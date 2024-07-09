package user

import (
	"errors"
	"go-learning/helpers/common"
)

var DummyUser = []User{
	{
		Username: "admin",
		Password: "admin",
		Role:     "super_admin",
	},
	{
		Username: "developer",
		Password: "developer",
		Role:     "passive_user",
	},
}

type User struct {
	Username string
	Password string
	Role     string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginRequest) ValidateLogin() (err error) {
	if common.CheckIsStringEmpty(l.Username) {
		return errors.New("username required")
	}

	if common.CheckIsStringEmpty(l.Password) {
		return errors.New("password required")
	}

	return
}

type LoginResponse struct {
	Token string `json:"token"`
}
