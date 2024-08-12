package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-learning/middlewares"
	"go-learning/utils/common"
	"go-learning/utils/constant"
	"go-learning/utils/email"
	"go-learning/utils/logger"
	"go-learning/utils/rabbitmq"
	"go-learning/utils/redis"
	"time"
)

type Service interface {
	LoginService(ctx *gin.Context, rabbitMqConn *rabbitmq.RabbitMQ, req LoginRequest) (result LoginResponse, err error)
	SignUpService(ctx *gin.Context, req SignUpRequest) (err error)
	GetEmailNotification(ctx *gin.Context) (resp string, err error)
}

type userService struct {
	repository      Repository
	emailRepository email.Repository
}

func NewService(repository Repository, emailRepository email.Repository) Service {
	return &userService{
		repository, emailRepository,
	}
}

func (service *userService) LoginService(ctx *gin.Context, rabbitMqConn *rabbitmq.RabbitMQ, req LoginRequest) (result LoginResponse, err error) {
	var (
		redisPermission []middlewares.RedisPermission
	)

	// user section
	user, err := service.repository.Login(ctx, req)
	if err != nil {
		return
	}

	if common.IsEmptyField(user.ID) {
		err = errors.New("invalid account")
		return
	}

	matches := common.CheckPassword(user.Password, req.Password)
	if !matches {
		err = errors.New("wrong username or password")
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	// get permission from user data
	permissions, err := service.repository.GetListPermissionByRoleId(ctx, user)
	if err != nil {
		err = errors.New("failed get permissions")
		return LoginResponse{}, err
	}

	for _, permission := range permissions {
		redisPermission = append(redisPermission, middlewares.RedisPermission{
			AccessCode:  permission.AccessCode,
			AccessGrant: permission.GrantCode,
		})
	}

	// redis section
	redisSession := middlewares.RedisSession{
		UserId:     user.ID,
		Username:   user.Username,
		LoginAt:    time.Now(),
		RoleId:     user.RoleId,
		Permission: redisPermission,
		ExpiredAt:  time.Now().Add(time.Minute * 20),
	}

	redisSessionStr, err := json.Marshal(redisSession)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return result, err
	}

	// set value redis session
	if viper.GetString("app.mode") != constant.DevelopmentMode.String() {
		err = redis.RedisClient.Set(ctx, jwtToken, string(redisSessionStr), 0).Err()
		if err != nil {
			err = errors.New("failed set value redis")
			logger.ErrorWithCtx(ctx, nil, err)
			return
		}
	} else {
		middlewares.DummyRedis[jwtToken] = string(redisSessionStr)
	}

	ctx.Set("session", string(redisSessionStr))

	result.Token = jwtToken

	// send email & publish to rabbit mq
	if viper.GetString("app.mode") == constant.StagingMode.String() {
		emailStr, _ := service.GetEmailNotification(ctx)

		// publish to rabbitmq
		_ = rabbitMqConn.Publish(rabbitmq.MqConfig{
			QueueName: constant.EmailQueue,
			Messsage:  emailStr,
		})
	}

	return
}

func (service *userService) SignUpService(ctx *gin.Context, req SignUpRequest) (err error) {
	user, err := req.ConvertToModelForSignUp()
	if err != nil {
		return err
	}

	err = service.repository.SignUp(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) GetEmailNotification(ctx *gin.Context) (resp string, err error) {
	var (
		session middlewares.RedisSession
	)

	emailTemplate, err := service.emailRepository.GetEmailTemplate(ctx, constant.LoginEmailTemplate.String())
	if err != nil {
		err = errors.New("failed to get email template")
		return
	}

	device := common.GetDeviceInfo()
	_ = json.Unmarshal([]byte(ctx.GetString("session")), &session)

	userNotification := email.EmailNotif{
		Sender:  viper.GetString("notification.email.sender"),
		Subject: "Login Activity",
		Receiver: []string{
			"arifyuniarto88@gmail.com",
		},
		Message: emailTemplate.Template,
		Data: map[string]string{
			"Device":    device.Name,
			"Name":      session.Username,
			"IPAddress": device.IPAddress,
			"Location":  device.Location,
			"LoginTime": time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	message, err := json.Marshal(&userNotification)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string(message), nil
}
