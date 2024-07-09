package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/configs"
	"go-learning/modules/car"
	"go-learning/modules/user"
)

func main() {
	configs.InitiateConfiguration()
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	car.Initiator(router)
	user.Initiator(router)

	router.Run(":8080")
}
