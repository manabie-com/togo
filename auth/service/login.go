package service

import (
	"context"
	"github.com/manabie-com/togo/auth/dto"
	"github.com/manabie-com/togo/auth/model"
	"github.com/manabie-com/togo/shared"
	tokenprovider "github.com/manabie-com/togo/token_provider"
)

type FindUserStorage interface {
	FindUserById(ctx context.Context, userId string) (*model.User, error)
	FindUserByLoginId(ctx context.Context, userId string) (*model.User, error)
}

type authService struct {
	store      FindUserStorage
	tokProvier tokenprovider.Provider
}

func NewAuthService(store FindUserStorage, tokProvier tokenprovider.Provider) *authService {
	return &authService{
		store:      store,
		tokProvier: tokProvier,
	}
}

func (s *authService) Login(ctx context.Context, data *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.store.FindUserByLoginId(ctx, data.LoginId)

	if err != nil {
		return nil, shared.ErrCannotGetEntity(model.EntityName, err)
	}

	//TODO: you can check user state here to handler case user has been block
	if user.Status == 0 {
		return nil, model.ErrUserHasBeenBlock
	}

	if user.Password != data.Password {
		return nil, model.ErrIdOrPasswordInvalid
	}

	atok, err := s.tokProvier.GenAccessToken(user, 15000)
	if err != nil {
		return nil, shared.ErrInternal(err)
	}

	return &dto.LoginResponse{
		AccessToken: *atok,
	}, nil
}
