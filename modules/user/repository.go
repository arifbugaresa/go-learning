package user

import (
	"database/sql"
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
	sqlStmt := "SELECT id, password FROM " + constant.UserTableName.String() + " WHERE username = $1"

	params := []interface{}{
		user.Username,
	}

	err = r.db.QueryRow(sqlStmt, params...).Scan(&result.ID, &result.Password)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}

	return result, nil
}

func (r *userRepository) SignUp(user User) (err error) {
	sqlStmt := "INSERT INTO " + constant.UserTableName.String() + " (username, full_name, password) VALUES ($1, $2, $3)"

	params := []interface{}{
		user.Username,
		user.Username,
		user.Password,
	}

	_, err = r.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}
