package user

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
	"time"
)

type userService struct {
	repo Repository
	cache Cache
}

func NewService(repo Repository, cache Cache) *userService {
	return &userService{repo, cache}
}

func (s *userService) GetUser(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil
}

func (s *userService) IsOverLimitTask(ctx context.Context, userId string) bool {
	return false
}

func (s *userService) IncreaseTaskTimesPerDuration(ctx context.Context, userId string, duration time.Duration) (int, error) {
	return 0, nil
}

