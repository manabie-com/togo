package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

var CACHE_TOKEN_KEY = func(token string) string { return fmt.Sprintf("login_token:%s", token) }

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (r *RedisCache) AddToken(ctx context.Context, token string, exp time.Duration) error {
	return r.client.Set(ctx, CACHE_TOKEN_KEY(token), token, exp).Err()
}

func (r *RedisCache) ValidateToken(ctx context.Context, token string) error {
	_, err := r.client.Get(ctx, CACHE_TOKEN_KEY(token)).Result()
	return err
}
