package config

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var RC *redis.Client

func Redis() {
	addr := os.Getenv("REDIS_URL")

	RC = redis.NewClient(&redis.Options{Addr: addr, Password: "", DB: 0})

	_, err := RC.Ping().Result()

	if err != nil {
		log.Panic(err)
	}

	log.Println("Redis Connected")
}
