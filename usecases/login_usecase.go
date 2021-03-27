package usecases

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/pkg/core"
)

var (
	ErrorInvalidUsernameOrPassword = errors.New("invalid username or password")
)

type (
	LoginInput struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginUseCase interface {
		Execute(ctx context.Context, request *LoginInput) (string, error)
	}

	loginInteractor struct {
		userRepo domains.UserRepository
		auth     core.AppAuthenticator
	}
)

func NewLoginUseCase(userRepo domains.UserRepository, auth core.AppAuthenticator) LoginUseCase {
	return loginInteractor{
		userRepo: userRepo,
		auth:     auth,
	}
}

// Execute login with dependencies
func (i loginInteractor) Execute(ctx context.Context, request *LoginInput) (string, error) {
	user, err := i.userRepo.VerifyUser(ctx, &domains.LoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		if err == domains.ErrorNotFound {
			return "", ErrorInvalidUsernameOrPassword
		}
		return "", err
	}

	authToken, err := i.auth.CreateToken(user.Id)
	if err != nil {
		return "", err
	}

	return authToken, nil
}
