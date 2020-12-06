package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/internal/storages/user"
)

//Service interface
type Service interface {
	Auth(ctx context.Context, user user.User, password string) error
	CreateToken(ctx context.Context, user user.User) string
}
