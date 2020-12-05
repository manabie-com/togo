package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

//Repository interface
type Repository interface {
	GetUser(ctx context.Context, id int64) (user.User, error)
}

//UseCase interface
type Service interface {
	Auth(ctx context.Context, id string) bool
}
