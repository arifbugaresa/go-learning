package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go-learning/middlewares"
	"go-learning/utils/common"
	"go-learning/utils/logger"
	"time"
)

type Service interface {
	LoginService(ctx *gin.Context) (result LoginResponse, err error)
	SignUpService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
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

	middlewares.DummyRedis[jwtToken] = string(redisSessionStr)

	result.Token = jwtToken

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
