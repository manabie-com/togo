package redis

import (
	"context"
	userUsecase "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/go-redis/redis"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewCache(redisClient *redis.Client) userUsecase.Cache {
	return &RedisClient{client: redisClient}
}

func (rdc *RedisClient) CheckLimit(ctx context.Context, userKey string, limitTime int) (bool, error) {
	return false, nil
}

func (rdc *RedisClient) IncreaseTask(ctx context.Context, userId string, duration time.Duration) (int, error) {
	return 0, nil
}
