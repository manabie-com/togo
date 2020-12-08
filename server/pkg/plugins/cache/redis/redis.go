package redis

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

//NewRedisClient new Redis Client for interacting redis
func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.address"),
		MaxRetries: viper.GetInt("redis.max_retries"),
	})

	if redisClient.Ping().Err() != nil {
		panic(" Connection Redis Error")
	}

	return redisClient
}
