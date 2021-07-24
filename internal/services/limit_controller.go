package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type limitController interface {
	ReachLimit(ctx context.Context, userID, date string, limitPerDay int64) (fallback func(), err error)
}

func newLimitController() *redisCache {
	return &redisCache{client: redis.NewClient(&redis.Options{
		Addr: redisHost,
	})}
}

type redisCache struct {
	client *redis.Client
}

func (c *redisCache) ReachLimit(ctx context.Context, userID, date string, limitPerDay int64) (fallback func(), err error) {
	limitCreatedUserTaskPerDayCacheKey := createLimitCacheKey(userID, date)
	val, err := c.client.Incr(ctx, limitCreatedUserTaskPerDayCacheKey).Result()
	if err != nil {
		return nil, err
	}
	if val > limitPerDay {
		//try to fallback, so the value should be the limit number
		if decrErr := c.client.Decr(ctx, limitCreatedUserTaskPerDayCacheKey).Err(); decrErr != nil {
			log.Println(decrErr)
		}
		return nil, newError(http.StatusTooManyRequests, limitErrorMsg)
	}
	if cleanupErr := c.client.ExpireAt(ctx, limitCreatedUserTaskPerDayCacheKey, endOfToday); cleanupErr != nil {
		log.Println(cleanupErr)
	}
	return func() {
		if decrErr := c.client.Decr(ctx, limitCreatedUserTaskPerDayCacheKey).Err(); decrErr != nil {
			log.Println(decrErr)
		}
		if cleanupErr := c.client.ExpireAt(ctx, limitCreatedUserTaskPerDayCacheKey, endOfToday); cleanupErr != nil {
			log.Println(cleanupErr)
		}
	}, nil
}

func createLimitCacheKey(userID, date string) string {
	return fmt.Sprintf("%s:%s:%s", limitCacheKeyPrefix, date, userID)
}
