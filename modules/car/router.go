package car

import (
	"github.com/gin-gonic/gin"
	"go-learning/middlewares"
	"go-learning/utils/common"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/cars")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("", CreateCarRouter)
		api.GET("", GetAllCarRouter)
		api.GET("/:id", GetCarRouter)
		api.PUT("/:id", UpdateCarRouter)
		api.DELETE("/:id", DeleteCarRouter)
	}
}

func CreateCarRouter(ctx *gin.Context) {
	var (
		carRepo = NewRepository()
		carSrv  = NewService(carRepo)
	)

	_, err := carSrv.CreateCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added car data")
}

func GetAllCarRouter(ctx *gin.Context) {
	var (
		carRepo = NewRepository()
		carSrv  = NewService(carRepo)
	)

	cars, err := carSrv.GetAllCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all car data", cars)
}

func GetCarRouter(ctx *gin.Context) {
	var (
		carRepo = NewRepository()
		carSrv  = NewService(carRepo)
	)

	cars, err := carSrv.GetCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get car data", cars)
}

func DeleteCarRouter(ctx *gin.Context) {
	var (
		carRepo = NewRepository()
		carSrv  = NewService(carRepo)
	)

	err := carSrv.DeleteCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete car data")
}

func UpdateCarRouter(ctx *gin.Context) {
	var (
		carRepo = NewRepository()
		carSrv  = NewService(carRepo)
	)

	err := carSrv.UpdateCarService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update car data")
}
