package redis_driver

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// RedisConfiguration ...
type RedisConfiguration struct {
	Addr string
}

// RedisDriver ...
type RedisDriver struct{}

// RedisClient ...
var RedisClient *redis.Client

// Setup ...
func (driver RedisDriver) Setup(config RedisConfiguration) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.Addr,
	})
	if err := RedisClient.Ping().Err(); err != nil {
		return errors.Wrapf(err, "Error while checking redis connection")
	}
	log.Println("Redis connection successful")

	return nil
}
