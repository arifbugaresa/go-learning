package user

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"go-learning/helpers/constant"
)

type Repository interface {
	Login(user LoginRequest) (result User, err error)
	SignUp(user User) (err error)
}

type userRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) Login(user LoginRequest) (result User, err error) {
	conn := goqu.New(constant.PostgresDialect.String(), r.db)
	dialect := conn.From(constant.UserTableName.String()).
		Select(goqu.C("id")).
		Where(
			goqu.I("username").Eq(user.Username),
			goqu.I("password").Eq(user.Password),
		)

	_, err = dialect.ScanStruct(&result)
	if err != nil {
		return
	}

	return
}

func (r *userRepository) SignUp(user User) (err error) {
	conn := goqu.New(constant.PostgresDialect.String(), r.db)
	dataset := conn.Insert(constant.UserTableName.String()).Rows(
		goqu.Record{
			"username":  user.Username,
			"full_name": user.Username,
			"password":  user.Password,
		},
	)

	_, err = dataset.Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}
