package cache

import "github.com/go-redis/redis/v8"

var (
	RedisClient *redis.Client
)

const (
	RedisHost = "localhost"
	RedisPort = 6379
)
