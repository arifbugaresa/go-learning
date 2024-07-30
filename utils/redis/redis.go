package redis

import (
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Initiator() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
