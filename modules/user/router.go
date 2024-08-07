package user

import (
	"github.com/gin-gonic/gin"
	"go-learning/databases/connection"
	"go-learning/middlewares"
	"go-learning/modules/notification"
	"go-learning/utils/common"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/users")
	api.Use(middlewares.Logging())
	{
		api.POST("/login", Login)
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
func Login(ctx *gin.Context) {
	var (
		userRepo  = NewRepository(connection.DBConnections)
		emailRepo = notification.NewRepository(connection.DBConnections)
		userSrv   = NewService(userRepo, emailRepo)
	)

	token, err := userSrv.LoginService(ctx)
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
		emailRepo = notification.NewRepository(connection.DBConnections)
		userSrv   = NewService(userRepo, emailRepo)
	)

	err := userSrv.SignUpService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "awesome, successfully create user")
}
