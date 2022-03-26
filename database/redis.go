package database

import (
	"fmt"
	"log"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/luongdn/togo/config"
)

var Rdb *redis.Client

func ConnectRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Host, config.Config.Redis.Port),
		Password: "",
		DB:       0,
	})
	log.Println("Redis connected")
}

func NewMockRedisClient(server *miniredis.Miniredis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: server.Addr(),
	})
}
