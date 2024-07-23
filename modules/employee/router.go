package employee

import (
	"github.com/gin-gonic/gin"
	"go-learning/databases/connection"
	"go-learning/helpers/common"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/employees")
	{
		api.GET("", ListEmployee)
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
