package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-learning/middlewares"
	"go-learning/utils/common"
	"go-learning/utils/email"
	"go-learning/utils/rabbitmq"
)

func Initiator(router *gin.Engine, rabbitMqConn *rabbitmq.RabbitMQ, dbConnection *sql.DB) {
	var (
		userRepo  = NewRepository(dbConnection)
		emailRepo = email.NewRepository(dbConnection)
		userSrv   = NewService(userRepo, emailRepo)
	)

	api := router.Group("/api/users")
	api.Use(middlewares.Logging())
	{
		api.POST("/login", func(c *gin.Context) {
			Login(c, rabbitMqConn, userSrv)
		})
		api.POST("/signup", func(c *gin.Context) {
			SignUp(c, userSrv)
		})
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
	var (
		req LoginRequest
	)

	// validation request section
	err := ctx.ShouldBind(&req)
	if err != nil {
		return
	}

	err = req.ValidateLogin()
	if err != nil {
		return
	}

	token, err := userSrv.LoginService(ctx, rabbitMqConn, req)
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
func SignUp(ctx *gin.Context, userSrv Service) {
	var req SignUpRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		common.GenerateErrorResponse(ctx, "invalid request body")
		return
	}

	err = req.ValidateSignUp()
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err = userSrv.SignUpService(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "awesome, successfully create user")
}
