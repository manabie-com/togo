package services

import (
	"context"
	"net/http"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	errPkg "github.com/manabie-com/togo/pkg/errors"
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
		return "", errPkg.NewCustomError("failed to validate user", http.StatusInternalServerError)
	}
	if userReq.Password != user.Password {
		return "", errPkg.NewCustomError("password is wrong", http.StatusInternalServerError)
	}

	token, err := httpPkg.CreateToken(user.ID)
	if err != nil {
		return "", errPkg.NewCustomError("failed to create token", http.StatusInternalServerError)
	}
	return token, nil
}
