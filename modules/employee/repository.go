package employee

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"go-learning/utils/common"
	"go-learning/utils/constant"
	"go-learning/utils/database"
)

type Repository interface {
	GetAllEmployee(ctx *gin.Context, request GetEmployeeRequest) (result []Employee, total int64, err error)
}

type empRepository struct {
	db *sql.DB
}

func NewRepository(dbParam *sql.DB) Repository {
	return &empRepository{
		db: dbParam,
	}
}

func (r *empRepository) GetAllEmployee(ctx *gin.Context, request GetEmployeeRequest) (result []Employee, total int64, err error) {
	var (
		sortField = "id"
		sortOrder = "ASC"
	)

	conn := goqu.New(constant.PostgresDialect.String(), r.db)
	dataset := conn.From(constant.EmployeeTableName.String()).
		Select(
			goqu.C("id"),
			goqu.C("full_name"),
			goqu.C("email"),
			goqu.C("age"),
			goqu.C("division"),
			goqu.C("modified_by"),
			goqu.C("modified_at"),
			goqu.C("created_at"),
			goqu.C("created_by"),
		)

	if !common.IsEmptyField(request.SearchBy.FullName) {
		dataset = dataset.Where(
			goqu.I("full_name").ILike("%" + *request.SearchBy.FullName + "%"),
		)
	}

	if (!common.IsEmptyField(request.SortField)) && (!common.IsEmptyField(request.SortOrder)) {
		sortField = *request.SortField
		sortOrder = *request.SortOrder
	}

	dataset, total, err = database.BuildDatasetPaginationWithTotalData(dataset, request.Page, request.Limit, sortField, sortOrder)

	err = dataset.ScanStructs(&result)
	if err != nil {
		return
	}

	return
}
