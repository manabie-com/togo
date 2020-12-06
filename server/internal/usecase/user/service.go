package user

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

type userService struct {
	repo Repository
}

func NewService(repo Repository) *userService {
	return &userService{repo}
}

func (s *userService) GetUser(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil
}
