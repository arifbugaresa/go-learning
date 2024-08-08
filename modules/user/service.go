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
	"go-learning/utils/redis"
	"time"
)

type Service interface {
	LoginService(ctx *gin.Context) (result LoginResponse, err error)
	SignUpService(ctx *gin.Context) (err error)
	SendEmailNotification(ctx *gin.Context) (err error)
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

func (service *userService) LoginService(ctx *gin.Context) (result LoginResponse, err error) {
	var (
		userReq         LoginRequest
		redisPermission []middlewares.RedisPermission
	)

	// validation request section
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateLogin()
	if err != nil {
		return
	}

	// user section
	user, err := service.repository.Login(ctx, userReq)
	if err != nil {
		return
	}

	if common.IsEmptyField(user.ID) {
		err = errors.New("invalid account")
		return
	}

	matches := common.CheckPassword(user.Password, userReq.Password)
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
		return
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
		ExpiredAt:  time.Now().Add(time.Minute * 3),
	}

	redisSessionStr, err := json.Marshal(redisSession)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return result, err
	}

	// select mode app
	if viper.GetString("app.mode") == "development" {
		middlewares.DummyRedis[jwtToken] = string(redisSessionStr)
	} else {
		err = redis.RedisClient.Set(ctx, jwtToken, string(redisSessionStr), 0).Err()
		if err != nil {
			err = errors.New("failed set value redis")
			logger.ErrorWithCtx(ctx, nil, err)
			return
		}
	}

	ctx.Set("session", string(redisSessionStr))

	result.Token = jwtToken

	if viper.GetString("app.mode") == "staging" {
		err = service.SendEmailNotification(ctx)
		if err != nil {
			return
		}
	}

	return
}

func (service *userService) SignUpService(ctx *gin.Context) (err error) {
	var userReq SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}

	err = userReq.ValidateSignUp()
	if err != nil {
		return err
	}

	user, err := userReq.ConvertToModelForSignUp()
	if err != nil {
		return err
	}

	err = service.repository.SignUp(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) SendEmailNotification(ctx *gin.Context) (err error) {
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

	userNotification.SendEmail(ctx)

	return
}
