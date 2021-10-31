package redis

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/redis_driver"
)

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

// TokenBlackListCacheKey ...
func TokenBlackListCacheKey(key string) string {
	return ":token:BlackList:" + key
}
