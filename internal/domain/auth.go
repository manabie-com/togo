package domain

import (
	"context"

	"github.com/perfectbuii/togo/common/errors"
	"github.com/perfectbuii/togo/internal/storages"
	"github.com/perfectbuii/togo/utils"
)

type AuthDomain interface {
	Login(ctx context.Context, u *storages.User) (string, error)
	Register(ctx context.Context, u *storages.User) error
}

type authDomain struct {
	userStore storages.UserStore
	jwtKey    string
}

func (d *authDomain) Login(ctx context.Context, u *storages.User) (string, error) {
	user, err := d.userStore.FindUser(ctx, u.ID)
	if err != nil {
		return "", errors.ErrUserDoesNotExist
	}
	if err := utils.VerifyPassword(u.Password, user.Password); err != nil {
		return "", err
	}

	return utils.CreateToken(user.ID, d.jwtKey)
}

func (d *authDomain) Register(ctx context.Context, u *storages.User) error {

	if _, err := d.userStore.FindUser(ctx, u.ID); err == nil {
		return errors.ErruserAlreadyExist
	}
	pwd, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwd
	u.MaxTodo = utils.GetNumRandomTask()
	return d.userStore.Create(ctx, u)
}

func NewAuthDomain(userStore storages.UserStore, jwtKey string) AuthDomain {
	return &authDomain{
		userStore: userStore,
		jwtKey:    jwtKey,
	}
}
