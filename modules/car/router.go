package car

import (
	"github.com/gin-gonic/gin"
	"go-learning/middlewares"
	"go-learning/utils/common"
)

func Initiator(router *gin.Engine) {
	var (
		carRepo = NewRepository()
		carSrv  = NewService(carRepo)
	)

	api := router.Group("/api/cars")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", func(c *gin.Context) {
			CreateCarRouter(c, carSrv)
		})
		api.GET("", func(c *gin.Context) {
			GetAllCarRouter(c, carSrv)
		})
		api.GET("/:id", func(c *gin.Context) {
			GetCarRouter(c, carSrv)
		})
		api.PUT("/:id", func(c *gin.Context) {
			UpdateCarRouter(c, carSrv)
		})
		api.DELETE("/:id", func(c *gin.Context) {
			DeleteCarRouter(c, carSrv)
		})
	}
}

func CreateCarRouter(ctx *gin.Context, carSrv Service) {
	_, err := carSrv.CreateCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added car data")
}

func GetAllCarRouter(ctx *gin.Context, carSrv Service) {
	cars, err := carSrv.GetAllCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all car data", cars)
}

func GetCarRouter(ctx *gin.Context, carSrv Service) {
	cars, err := carSrv.GetCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get car data", cars)
}

func DeleteCarRouter(ctx *gin.Context, carSrv Service) {
	err := carSrv.DeleteCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete car data")
}

func UpdateCarRouter(ctx *gin.Context, carSrv Service) {
	err := carSrv.UpdateCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update car data")
}
