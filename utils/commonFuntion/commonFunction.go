package commonFuntion

import (
	"fmt"
	"go-learning/utils/constanta"
)

func GetDBConnection() error {
	message := "success get db connections: " + constanta.DBPort

	fmt.Println(message)

	return nil
}

func GetRedisConnection() error {
	message := "success get redis connections: " + constanta.RedisPort

	fmt.Println(message)

	return nil
}

func GetElasticConnection() error {
	message := "success get elastic connections: " + constanta.ElasticPort

	fmt.Println(message)

	return nil
}
