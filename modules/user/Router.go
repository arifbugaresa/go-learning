package user

import (
	"github.com/gin-gonic/gin"
	"go-learning/helpers/common"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/login", LoginRouter)
	}
}

func LoginRouter(ctx *gin.Context) {
	var (
		userRepo = NewRepository()
		userSrv  = NewService(userRepo)
	)

	token, err := userSrv.LoginService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully login", token)
}
