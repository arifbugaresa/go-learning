package scheduler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learning/modules/car"
)

func Initiator() {
	scheduler := NewScheduler()

	scheduler.AddJob("@every 1s", LogAllCar())

	//scheduler.Start()
}

func LogAllCar() func() {
	return func() {
		ctx := gin.Context{}

		carRepo := car.NewRepository()
		carServ := car.NewService(carRepo)

		cars, _ := carServ.GetAllCarService(&ctx)

		fmt.Println("cars: ", cars)
	}
}
