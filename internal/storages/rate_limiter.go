package storages

import (
	"context"
)

//go:generate mockgen -source=rate_limiter.go -destination=./mocks/rate_limiter_mock.go

// RateLimiter use for rate limit config and auth user request
type RateLimiter interface {
	Allow(ctx context.Context, userID string) (bool, error)
}
