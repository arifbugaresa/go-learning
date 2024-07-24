package employee

import (
	"github.com/gin-gonic/gin"
	"go-learning/databases/connection"
	"go-learning/middlewares"
	"go-learning/utils/common"
)

func Initiator(router *gin.Engine) {
	var (
		GetListPermission = map[string]string{"employee": "r"}
	)

	api := router.Group("/api/employees")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.GET("", middlewares.Permission(GetListPermission), ListEmployee)
	}
}

func ListEmployee(ctx *gin.Context) {
	var (
		empRepo = NewRepository(connection.DBConnections)
		empSrv  = NewService(empRepo)
	)

	data, total, err := empSrv.GetListEmployee(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all employee data", total, data)
}
