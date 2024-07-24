package employee

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-learning/utils/logger"
)

type Service interface {
	GetListEmployee(ctx *gin.Context) (result []GetEmployeeResponse, total int64, err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) GetListEmployee(ctx *gin.Context) (result []GetEmployeeResponse, total int64, err error) {
	var (
		req GetEmployeeRequest
	)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		err = errors.New("failed to parse request body")
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	employees, total, err := service.repository.GetAllEmployee(req)
	if err != nil {
		return
	}

	// convert to response
	for _, employee := range employees {
		result = append(result, employee.ConvertModelToResponseForGetListEmployee())
	}

	return
}
