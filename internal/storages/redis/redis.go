package redis

import (
	"github.com/go-redis/redis/v8"
)


func init() {
	Init()
}

func Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}