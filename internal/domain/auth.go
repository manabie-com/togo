package domain

import (
	"context"
	"fmt"

	"github.com/manabie-com/togo/common/errors"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils"
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
	fmt.Println(user)
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
	return d.userStore.Create(ctx, u)
}

func NewAuthDomain(userStore storages.UserStore, jwtKey string) AuthDomain {
	return &authDomain{
		userStore: userStore,
		jwtKey:    jwtKey,
	}
}
