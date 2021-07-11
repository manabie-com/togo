package auth

import (
	"context"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/pkg/jwtprovider"
)

type service struct {
	userStorage storages.UserStorage
	jwtProvider jwtprovider.JWTProvider
}

type Service interface {
	Login(ctx context.Context, userID, pwd string) (string, error)
}

func NewAuthService(userStorage storages.UserStorage, jwtProvider jwtprovider.JWTProvider) Service {
	return &service{
		userStorage: userStorage,
		jwtProvider: jwtProvider,
	}
}

func (s *service) Login(ctx context.Context, userID, pwd string) (string, error) {
	user, err := s.userStorage.Find(ctx, userID, pwd)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", ErrWrongAccount
	}
	return s.jwtProvider.GenerateToken(map[string]interface{}{
		"user_id": userID,
	})
}
