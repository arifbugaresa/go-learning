package main

import (
	"github.com/gin-gonic/gin"
	"go-learning/configs"
	"go-learning/databases/connection"
	"go-learning/databases/migration"
	"go-learning/modules/car"
	"go-learning/modules/employee"
	"go-learning/modules/user"
	"go-learning/utils/logger"
	"go-learning/utils/rabbitmq"
	"go-learning/utils/redis"
	"go-learning/utils/scheduler"
	"go-learning/utils/swagger"
)

// @title Swagger Documentation
// @version 1.1.2
// @description This is documentation go_learning.
// @host localhost:8080
func main() {
	configs.Initiator()

	logger.Initiator()

	scheduler.Initiator()

	redis.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	migration.Initiator(connection.DBConnections)

	// initiate rabbitmq publisher
	rabbitMqConn := rabbitmq.Initiator()
	defer rabbitMqConn.Channel.Close()
	defer rabbitMqConn.Conn.Close()

	// initiate rabbitmq consumer
	_ = rabbitMqConn.Consume()

	// initiate router
	InitiateRouter(rabbitMqConn)
}

func InitiateRouter(rabbitMqConn *rabbitmq.RabbitMQ) {
	router := gin.Default()

	swagger.Initiator(router)

	car.Initiator(router)
	user.Initiator(router, rabbitMqConn)
	employee.Initiator(router)

	router.Run(":8080")
}
