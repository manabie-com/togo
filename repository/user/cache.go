package user

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type cacheRepository struct {
	getClient func(ctx context.Context) *redis.Client
}

func NewCacheRepository(getClient func(ctx context.Context) *redis.Client) CacheRepository {
	return &cacheRepository{
		getClient: getClient,
	}
}
