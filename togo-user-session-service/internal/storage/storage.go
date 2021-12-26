package storage

import (
	"context"
	"errors"

	"togo-user-session-service/internal/model"
)

var (
	ErrWrongUserNameOrPass = errors.New("wrong username or password")
	ErrTokenInvalid        = errors.New("token invalid")
)

type Storage interface {
	RegisterOrLogin(ctx context.Context, username, password string) (string, error)
	VerifyToken(ctx context.Context, token string) (*model.User, error)
}
