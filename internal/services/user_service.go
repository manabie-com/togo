package services

import (
	"context"
	"github.com/manabie-com/togo/internal/constants"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/sirupsen/logrus"
)

type userService struct {
	tokenProvider  helpers.TokenProvider
	userRepository repositories.UserRepository
}

type UserService interface {
	GetAuthToken(ctx context.Context, userId, password string) (string, error)
}

func NewUserService(injectedTokenProvider helpers.TokenProvider,
	injectedUserRepository repositories.UserRepository) UserService {
	return &userService{
		tokenProvider:  injectedTokenProvider,
		userRepository: injectedUserRepository,
	}
}

func (s *userService) GetAuthToken(ctx context.Context, userId, password string) (string, error) {
	if !s.userRepository.ValidateUser(ctx, userId, password) {
		logrus.Errorf("Validating User not match :%s", userId)
		return "", constants.ErrIncorrectUserIdOrPassword
	}

	token, err := s.tokenProvider.CreateToken(userId)
	if err != nil {
		logrus.Errorf("CreateToken error: %s", err.Error())
		return "", constants.ErrCreateToken
	}

	return token, nil
}
