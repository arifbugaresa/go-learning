package employee

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"go-learning/helpers/constant"
)

type Repository interface {
	GetAllEmployee() (result []Employee, err error)
}

type empRepository struct {
	db *sql.DB
}

func NewRepository(dbParam *sql.DB) Repository {
	return &empRepository{
		db: dbParam,
	}
}

func (r *empRepository) GetAllEmployee() (result []Employee, err error) {
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

	err = dataset.Prepared(true).ScanStructs(&result)
	if err != nil {
		fmt.Println(err)
	}

	return
}
