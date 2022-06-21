package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func StartRedis() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", RedisHost, RedisPort),
	})

	_, err := c.Ping(context.Background()).Result()

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (r *Redis) GetInt(key string) (int, error) {
	var limit int
	v, err := r.client.Get(context.Background(),key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return limit, err
		}
		return limit, nil
	}
	limit, err = strconv.Atoi(v)
	if err != nil {
		return limit, fmt.Errorf("invalid integer: %v", err)
	}
	return limit, nil
}

func (r *Redis) SetExpire(key string, value int, expire time.Duration) error {
	result, err := r.client.SetEX(context.Background(), key, value, expire).Result()
	if result == "OK" {
		return nil
	}
	return err
}