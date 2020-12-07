package auth

import (
	"context"
)

type authService struct {
	repo Repository
}

//NewService create new service
func NewService(repo Repository) Service {
	return &authService{repo}
}

func (s *authService) Auth(ctx context.Context, userName, password string) (int64, error) {
	return 0, nil
}

