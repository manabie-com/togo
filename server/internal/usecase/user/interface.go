package user

import (
	"context"
	"time"
)

type Cache interface {
	IsOverLimit(ctx context.Context, userKey string, limitTime int) (bool, error)
	IncreaseTask(ctx context.Context, userId string, duration time.Duration) (int, error)
}

type Service interface {
	IsOverLimitTask(ctx context.Context, userId string, limit int) (bool, error)
	IncreaseTaskTimesPerDuration(ctx context.Context, userId string, duration time.Duration) (int, error)
}
