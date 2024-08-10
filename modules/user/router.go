package user

import (
	"github.com/gin-gonic/gin"
	"go-learning/databases/connection"
	"go-learning/middlewares"
	"go-learning/utils/common"
	"go-learning/utils/email"
	"go-learning/utils/rabbitmq"
)

func Initiator(router *gin.Engine, rabbitMqConn *rabbitmq.RabbitMQ) {
	var (
		userRepo  = NewRepository(connection.DBConnections)
		emailRepo = email.NewRepository(connection.DBConnections)
		userSrv   = NewService(userRepo, emailRepo)
	)

	api := router.Group("/api/users")
	api.Use(middlewares.Logging())
	{
		api.POST("/login", func(c *gin.Context) {
			Login(c, rabbitMqConn, userSrv)
		})
		api.POST("/signup", SignUp)
	}
}

// Login godoc
// @Tags User
// @Summary Login
// @Description	This endpoint is used for user login
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Request"
// @Success 200 {object} common.APIResponse{data=LoginResponse} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/users/login [post]
func Login(ctx *gin.Context, rabbitMqConn *rabbitmq.RabbitMQ, userSrv Service) {
	token, err := userSrv.LoginService(ctx, rabbitMqConn)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully login", token)
}

// SignUp godoc
// @Tags User
// @Summary Sign Up
// @Description	This endpoint is used for user sign up
// @Accept json
// @Produce json
// @Param login body SignUpRequest true "Request"
// @Success 200 {object} common.APIResponse "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/users/signup [post]
func SignUp(ctx *gin.Context) {
	var (
		userRepo  = NewRepository(connection.DBConnections)
		emailRepo = email.NewRepository(connection.DBConnections)
		userSrv   = NewService(userRepo, emailRepo)
	)

	err := userSrv.SignUpService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "awesome, successfully create user")
}
