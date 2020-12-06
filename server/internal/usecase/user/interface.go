package user

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
	"time"
)

//Repository interface
type Repository interface {
	GetUser(ctx context.Context, id int64) (user.User, error)
}

type Cache interface {
	CheckLimit(ctx context.Context, userKey string, limitTime int) bool
	IncreaseTask(ctx context.Context, userId string, time time.Duration) (int, error)
}

type Service interface {
	GetUser(ctx context.Context, id string) (user.User, error)
	IsOverLimitTask(ctx context.Context, userId string) bool
	IncreaseTaskTimesPerDuration(ctx context.Context, userId string, duration time.Duration) (int, error)
}
