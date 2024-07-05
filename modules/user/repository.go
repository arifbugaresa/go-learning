package user

type Repository interface {
	GetMessage(username string) (message string, err error)
}

type userRepository struct{}

func NewUserRepository() Repository {
	return new(userRepository)
}

func (r *userRepository) GetMessage(username string) (message string, err error) {
	greeting := "Hello "

	return greeting + username, nil
}
