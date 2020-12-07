package user

import (
	"context"
	"time"
)

type userService struct {
	cache Cache
}

func NewService(cache Cache) *userService {
	return &userService{cache}
}

func (s *userService) IsOverLimitTask(ctx context.Context, userId string, limit int) (bool, error) {
	return false, nil
}

func (s *userService) IncreaseTaskTimesPerDuration(ctx context.Context, userId string, duration time.Duration) (int, error) {
	return 0, nil
}

