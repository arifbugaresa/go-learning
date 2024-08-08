package email

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"go-learning/utils/constant"
	"go-learning/utils/logger"
)

type Repository interface {
	GetEmailTemplate(ctx *gin.Context, code string) (EmailTemplate, error)
}

type emailRepository struct {
	db *sql.DB
}

func NewRepository(dbParam *sql.DB) Repository {
	return &emailRepository{
		db: dbParam,
	}
}

func (r *emailRepository) GetEmailTemplate(ctx *gin.Context, code string) (res EmailTemplate, err error) {
	conn := goqu.New(constant.PostgresDialect.String(), r.db)
	dialect := conn.From(constant.EmailTemplateTableName.String()).
		Select(
			goqu.C("id"),
			goqu.C("code"),
			goqu.C("name"),
			goqu.C("template"),
		).
		Where(
			goqu.I("code").Eq(code),
		)

	_, err = dialect.ScanStruct(&res)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	return
}
