package car

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Service interface {
	CreateCarService(ctx *gin.Context) (result []Car, err error)
	GetAllCarService(ctx *gin.Context) (result []Car, err error)
	GetCarService(ctx *gin.Context) (result Car, err error)
	UpdateCarService(ctx *gin.Context) (err error)
	DeleteCarService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) CreateCarService(ctx *gin.Context) (result []Car, err error) {
	var newCar Car
	err = ctx.ShouldBind(&newCar)
	if err != nil {
		return
	}

	result, err = service.repository.CreateCarRepository(newCar)
	if err != nil {
		err = errors.New("failed to add new car")
		return nil, err
	}

	return
}

func (service *userService) GetAllCarService(ctx *gin.Context) (result []Car, err error) {
	return service.repository.GetAllCarRepository()
}

func (service *userService) GetCarService(ctx *gin.Context) (result Car, err error) {
	var (
		car Car
		id  = ctx.Param("id")
	)

	car.Id, err = strconv.Atoi(id)
	if err != nil {
		return
	}

	return service.repository.GetCarRepository(car)
}

func (service *userService) DeleteCarService(ctx *gin.Context) (err error) {
	var (
		car Car
		id  = ctx.Param("id")
	)

	car.Id, err = strconv.Atoi(id)
	if err != nil {
		return
	}

	return service.repository.DeleteCarRepository(car)
}

func (service *userService) UpdateCarService(ctx *gin.Context) (err error) {
	var (
		car Car
		id  = ctx.Param("id")
	)

	err = ctx.ShouldBind(&car)
	if err != nil {
		return
	}

	car.Id, err = strconv.Atoi(id)
	if err != nil {
		return
	}

	return service.repository.UpdateCarRepository(car)
}
