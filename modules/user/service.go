package user

type Service interface {
	GetMessage() string
}

type userService struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetMessage() string {
	username := "arifbugaresa"

	message, err := s.repo.GetMessage(username)
	if err != nil {
		panic(err)
	}

	return message
}
