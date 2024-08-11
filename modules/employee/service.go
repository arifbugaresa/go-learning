package employee

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListEmployee(ctx *gin.Context, req GetEmployeeRequest) (result []GetEmployeeResponse, total int64, err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) GetListEmployee(ctx *gin.Context, req GetEmployeeRequest) (result []GetEmployeeResponse, total int64, err error) {
	employees, total, err := service.repository.GetAllEmployee(ctx, req)
	if err != nil {
		return
	}

	// convert to response
	for _, employee := range employees {
		result = append(result, employee.ConvertModelToResponseForGetListEmployee())
	}

	return
}
