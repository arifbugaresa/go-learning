package employee

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-learning/middlewares"
	"go-learning/utils/common"
)

func Initiator(router *gin.Engine, dbConnection *sql.DB) {
	var (
		GetListPermission = map[string]string{"employee": "r"}
	)

	var (
		empRepo = NewRepository(dbConnection)
		empSrv  = NewService(empRepo)
	)

	api := router.Group("/api/employees")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.GET("", middlewares.Permission(GetListPermission), func(c *gin.Context) {
			ListEmployee(c, empSrv)
		})
	}
}

// ListEmployee godoc
// @Tags Employee
// @Summary Get List Employee
// @Description	This endpoint is used for get all employee
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse{data=GetEmployeeResponse} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/employees [get]
func ListEmployee(ctx *gin.Context, empSrv Service) {
	var req GetEmployeeRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		common.GenerateErrorResponse(ctx, "failed to parse request body")
		return
	}

	data, total, err := empSrv.GetListEmployee(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all employee data", total, data)
}
