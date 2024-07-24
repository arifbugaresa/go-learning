package user

import (
	"errors"
	"fmt"
	"go-learning/utils/common"
	"regexp"
)

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	RoleId   int64  `db:"role_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginRequest) ValidateLogin() (err error) {
	if common.IsEmptyField(l.Username) {
		return errors.New("username required")
	}

	if common.IsEmptyField(l.Password) {
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
	if common.IsEmptyField(s.Username) {
		return errors.New("username required")
	}

	if common.IsEmptyField(s.Password) {

		return errors.New("password required")
	}

	if common.IsEmptyField(s.ReTypePassword) {
		return errors.New("retype password required")
	}

	if s.ReTypePassword != s.Password {
		return errors.New("password mismatch!")
	}

	re := regexp.MustCompile(fmt.Sprintf(`^(.{8,})$`))
	if !re.MatchString(s.Password) {
		return errors.New("please make sure that the password contains at least 8 character")
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
