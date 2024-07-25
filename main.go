package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/configs"
	"go-learning/databases/connection"
	"go-learning/databases/migration"
	"go-learning/modules/car"
	"go-learning/modules/user"
)

func main() {
	configs.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	migration.Initiator(connection.DBConnections)

	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	car.Initiator(router)
	user.Initiator(router)

	router.Run(":8080")
}
