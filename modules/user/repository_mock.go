package user

import "github.com/stretchr/testify/mock"

type userRepositoryMock struct {
	mock.Mock
}

func (r *userRepositoryMock) GetMessage(username string) (message string, err error) {
	args := r.Called(username)

	return args.String(0), args.Error(1)
}
