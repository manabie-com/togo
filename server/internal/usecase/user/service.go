package user

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
	"time"
)

type userService struct {
	cache Cache
}

func NewService(cache Cache) *userService {
	return &userService{cache}
}

func (s *userService) IsOverLimitTask(ctx context.Context, userId string, limit int) (bool, error) {
	if userId == "" || limit == 0 {
		return false, define.FailedValidation
	}
	return s.cache.IsOverLimit(ctx, userId, limit)
}

func (s *userService) IncreaseTaskTimesPerDuration(ctx context.Context, userId string, duration time.Duration) (int, error) {
	if userId == "" || duration == 0 {
		return 0, define.FailedValidation
	}
	return s.cache.IncreaseTask(ctx, userId, duration)
}

