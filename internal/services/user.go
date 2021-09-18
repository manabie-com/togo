package services

import (
	"context"
	"errors"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	httpPkg "github.com/manabie-com/togo/pkg/http"
)

type UserService interface {
	GetAuthToken(ctx context.Context, user models.User) (string, error)
}

type userService struct {
	repo *repositories.Repository
}

func newUserService(repo *repositories.Repository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAuthToken(ctx context.Context, userReq models.User) (string, error) {
	user, err := s.repo.UserRepository.ValidateUser(ctx, userReq.ID)
	if err != nil {
		return "", err
	}
	if userReq.Password != user.Password {
		return "", errors.New("password is wrong")
	}

	token, err := httpPkg.CreateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
