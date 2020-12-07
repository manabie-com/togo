package auth

import (
	"context"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"github.com/HoangVyDuong/togo/pkg/utils"
)

type authService struct {
	repo Repository
}

//NewService create new service
func NewService(repo Repository) Service {
	return &authService{repo}
}

func (s *authService) Auth(ctx context.Context, userName, password string) (uint64, error) {
	if userName == "" || password == "" {
		logger.Error("[AuthService][Auth] param invalid")
		return 0, define.FailedValidation
	}

	user, err := s.repo.GetUserByName(ctx, userName)
	logger.Debugf("[AuthService][Auth] user from repository: %v", user)
	if err != nil {
		return 0, err
	}

	if isValidPassword := utils.ValidatePassword(password, user.Password) == nil; !isValidPassword {
		logger.Error("[AuthService][Auth] Password is not valid")
		return 0, define.AccountNotAuthorized
	}

	return user.ID, nil
}

