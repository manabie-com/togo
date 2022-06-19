package cache

import "github.com/go-redis/redis"

var (
	RedisClient *redis.Client
)

const (
	RedisHost = "localhost"
	RedisPort = 6379
)
