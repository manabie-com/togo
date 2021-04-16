package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"

	"github.com/manabie-com/togo/internal/storages"
)

var (
	rateLimiter *RateLimiter
	once        sync.Once
)

func GetRateLimiter(host, port string, maxRequestPerHour int) storages.RateLimiter {
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", host, port),
		})
		limiter := redis_rate.NewLimiter(rdb)

		rateLimiter = &RateLimiter{
			maxRequestPerHour: maxRequestPerHour,
			limiter:           limiter,
		}
	})

	return rateLimiter
}

type RateLimiter struct {
	maxRequestPerHour int
	limiter           *redis_rate.Limiter
}

func (r *RateLimiter) Allow(ctx context.Context, userID string) (bool, error) {
	key := fmt.Sprintf("ratelimit:%s", userID)

	result, err := r.limiter.Allow(ctx, key, redis_rate.PerHour(r.maxRequestPerHour))
	if err != nil {
		return false, err
	}

	return result.RetryAfter == time.Duration(-1), nil
}
