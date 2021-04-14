package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/services/users"
)

type Service interface {
	login(ctx context.Context, username string, password string) (string, error)
	getSecretJWT() string
}

type UID struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type service struct {
	userRepo  users.UserRepo
	secretJWT string
}

func NewAuthService(
	userRepo users.UserRepo,
	secretJWT string,
) *service {
	return &service{
		userRepo:  userRepo,
		secretJWT: secretJWT,
	}
}

func (s *service) getSecretJWT() string {
	return s.secretJWT
}
