package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/modules/car"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	car.Initiator(router)

	router.Run(":8080")
}
