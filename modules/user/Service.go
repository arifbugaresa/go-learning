package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-learning/helpers/common"
	"go-learning/middlewares"
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
	var userReq LoginRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateLogin()
	if err != nil {
		return
	}

	user, err := service.repository.Login(userReq)
	if err != nil {
		return
	}

	if common.CheckIsNumericEmpty(user.ID) {
		err = errors.New("user not found")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	middlewares.DummyRedis[jwtToken] = middlewares.UserLoginRedis{
		UserId:    0,
		Username:  user.Username,
		LoginAt:   time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 1),
	}

	result.Token = jwtToken

	return
}

func (service *userService) SignUpService(ctx *gin.Context) (err error) {
	var userReq SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateSignUp()
	if err != nil {
		return err
	}

	err = service.repository.SignUp(userReq.ConvertToModelForSignUp())
	if err != nil {
		err = errors.New("failed sign up user")
		return
	}

	return nil
}
