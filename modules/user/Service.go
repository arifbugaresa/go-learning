package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-learning/middlewares"
	"time"
)

type Service interface {
	LoginService(ctx *gin.Context) (result LoginResponse, err error)
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

	user, err := service.repository.LoginRepository(userReq)
	if err != nil {
		err = errors.New("invalid username or password")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	middlewares.DummyRedis[jwtToken] = middlewares.UserLoginRedis{
		UserId:    0,
		Username:  user.Username,
		Role:      user.Role,
		LoginAt:   time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 1),
	}

	result.Token = jwtToken

	return
}
