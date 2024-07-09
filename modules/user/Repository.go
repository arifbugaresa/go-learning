package user

import "errors"

type Repository interface {
	LoginRepository(user LoginRequest) (result User, err error)
}

type userRepository struct{}

func NewRepository() Repository {
	return &userRepository{}
}

func (r *userRepository) LoginRepository(user LoginRequest) (result User, err error) {
	for _, item := range DummyUser {
		if item.Username == user.Username && item.Password == user.Password {
			return item, nil
		}
	}

	return result, errors.New("user not found")
}
