package redis

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/go-redis/redis"
	"time"
)

type redisClient struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *redisClient {
	return &redisClient{client: client}
}

func (rdc *redisClient) IsOverLimit(ctx context.Context, userKey string, limitTime int) (bool, error) {
	times, err := rdc.client.Get(userKey).Int()
	if err != nil && err != redis.Nil {
		logger.Errorf("[Cache][IsOverLimit] error: %s", err.Error())
		return false, define.CacheError
	}
	if times >= limitTime {
		return true, nil
	}
	return false, nil
}

func (rdc *redisClient) IncreaseTask(ctx context.Context, userKey string, duration time.Duration) (int, error) {
	times, err := rdc.client.Incr(userKey).Result()
	if err != nil {
		return 0, define.CacheError
	}
	if times == 1 {
		rdc.client.Expire(userKey, duration)
	}
	return 0, nil
}
