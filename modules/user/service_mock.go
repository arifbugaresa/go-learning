package user

import "github.com/stretchr/testify/mock"

type UserServiceMock struct {
	mock.Mock
}

func (s *UserServiceMock) NewUserServiceMock(repo Repository) Service {
	return NewUserService(
		&UserRepositoryMock{},
	)
}

func (s *UserServiceMock) GetMessage() string {
	args := s.Called()
	return args.String(0) + args.String(1)
}
