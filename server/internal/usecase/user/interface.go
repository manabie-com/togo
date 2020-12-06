package user

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

//Repository interface
type Repository interface {
	GetUser(ctx context.Context, id int64) (user.User, error)
}

type Service interface {
	GetUser(ctx context.Context, id string) (user.User, error)
}
