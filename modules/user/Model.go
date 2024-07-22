package user

import (
	"errors"
	"go-learning/helpers/common"
)

var DummyUser = []UserOld{
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

type UserOld struct {
	Username string
	Password string
	Role     string
}

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
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

type SignUpRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"re_type_password"`
}

func (s *SignUpRequest) ValidateSignUp() (err error) {
	if common.CheckIsStringEmpty(s.Username) {
		return errors.New("username required")
	}

	if common.CheckIsStringEmpty(s.Password) {

		return errors.New("password required")
	}

	if common.CheckIsStringEmpty(s.ReTypePassword) {
		return errors.New("retype password required")
	}

	if s.ReTypePassword != s.Password {
		return errors.New("password mismatch!")
	}

	return nil
}

func (s *SignUpRequest) ConvertToModelForSignUp() (user User, err error) {
	hashedPassword, err := common.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password failed")
		return
	}

	return User{
		Username: s.Username,
		Password: hashedPassword,
	}, nil
}

type SignUpResponse struct {
}
