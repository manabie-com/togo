package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type CustomRedis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type Redis struct {
	client CustomRedis
}

func (r Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	fmt.Println("Get ", r.client)
	return r.client.Get(ctx, key)
}

func (r Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}
