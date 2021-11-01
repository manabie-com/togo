package redis

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/redis_driver"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
)

func Init() {
	// Init redis
	redisHost := os.Getenv(settings.RedisHost) + ":" + os.Getenv(settings.RedisPort)
	redisDriver := redis_driver.RedisDriver{}
	err := redisDriver.Setup(redis_driver.RedisConfiguration{
		Addr: redisHost,
	})
	if err != nil {
		panic(fmt.Sprint("Error while setup redis driver: ", err))
	}
}

// GetItem ...
func GetItem(key string, src interface{}) (interface{}, error) {
	val, err := redis_driver.RedisClient.Get(key).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "Get cache has error")
	}

	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return nil, errors.Wrapf(err, "Extra cache has error")
	}

	return val, nil
}

// SetItem ...
func SetItem(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return errors.Wrapf(err, "Input value cache has error")
	}

	err = redis_driver.RedisClient.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return errors.Wrapf(err, "Set cache has error")
	}

	return nil
}

// DeleteItem ...
func DeleteItem(key string) error {
	err := redis_driver.RedisClient.Del(key).Err()
	if err != nil {
		return errors.Wrapf(err, "Delete cache has error")
	}

	return nil
}
