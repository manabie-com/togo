package internal

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"togo/config"
)

func NewRedis(conf *config.Config) (*redis.Client, error) {

	dbi, _ := strconv.Atoi(conf.RedisDB)

	rdb := redis.NewClient(&redis.Options{
		Addr: conf.RedisHost,
		DB:   dbi,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("rdb.Ping %w", err)
	}

	return rdb, nil
}
