package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

// Repository interface
type Repository interface {
	GetUserByName(ctx context.Context, name string) (user.User, error)
}

// Service interface
type Service interface {
	Auth(ctx context.Context, userName, password string) (userID int64, err error)
}
