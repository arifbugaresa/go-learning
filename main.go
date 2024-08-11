package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	// initiate file configuration
	configs.Initiator()

	// initiate logger
	logger.Initiator()

	// initiate scheduler
	scheduler.Initiator()

	// initiate redis
	redis.Initiator()

	// initiate database connection
	dbConnection, _ := connection.Initiator()
	defer dbConnection.Close()

	// initiate sql migration
	migration.Initiator(dbConnection)

	// initiate rabbitmq publisher
	rabbitMqConn := rabbitmq.Initiator()
	defer rabbitMqConn.Channel.Close()
	defer rabbitMqConn.Conn.Close()

	// initiate rabbitmq consumer
	_ = rabbitMqConn.Consume()

	// initiate router
	InitiateRouter(dbConnection, rabbitMqConn)
}

func InitiateRouter(dbConnection *sql.DB, rabbitMqConn *rabbitmq.RabbitMQ) {
	router := gin.Default()

	// initiate swagger docs
	swagger.Initiator(router)

	car.Initiator(router)
	user.Initiator(router, rabbitMqConn, dbConnection)
	employee.Initiator(router, dbConnection)

	router.Run(viper.GetString("app.port"))
}
