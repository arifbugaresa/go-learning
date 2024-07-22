package employee

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListEmployee(ctx *gin.Context) (result []GetEmployeeResponse, err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) GetListEmployee(ctx *gin.Context) (result []GetEmployeeResponse, err error) {
	var (
		req GetEmployeeRequest
	)

	err = ctx.ShouldBindJSON(&req.SearchAndFilter)
	if err != nil {
		return nil, err
	}

	employees, err := service.repository.GetAllEmployee(req)
	if err != nil {
		err = errors.New("failed get all employees")
		return
	}

	// convert to response
	for _, employee := range employees {
		result = append(result, employee.ConvertModelToResponseForGetListEmployee())
	}

	return
}
