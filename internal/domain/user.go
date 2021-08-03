package domain

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"togo/common"
	"togo/internal/entity"
)

type UserHandler interface {
	CreateUser(ctx context.Context, username string, password string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUser(ctx context.Context, id int32) (*entity.User, error)
}

type UserDomain struct {
	handler UserHandler
}

func NewUserDomain(handler UserHandler) *UserDomain {
	return &UserDomain{handler: handler}
}

func (u *UserDomain) CreateUser(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := u.handler.CreateUser(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserDomain) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := u.handler.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserDomain) Login(ctx context.Context, username string, password string) (*entity.User, error) {
	user, err := u.handler.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	userPass := []byte(password)
	dbPass := []byte(user.Password)

	if passErr := bcrypt.CompareHashAndPassword(dbPass, userPass); passErr != nil {
		return nil, common.ErrPasswordNotMatch
	}

	return user, nil
}

func (u *UserDomain) GetUser(ctx context.Context, id int32) (*entity.User, error) {
	user, err := u.handler.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
