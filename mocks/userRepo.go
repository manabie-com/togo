package mocks

import (
	"context"

	"github.com/lawtrann/togo"
)

type UserRepo struct {
	GetUserByNameFn  func(ctx context.Context, username string) (*togo.User, error)
	IsExceedPerDayFn func(ctx context.Context, u *togo.User) (bool, error)
}

func (ur *UserRepo) GetUserByName(ctx context.Context, username string) (*togo.User, error) {
	return ur.GetUserByNameFn(ctx, username)
}

func (ur *UserRepo) IsExceedPerDay(ctx context.Context, u *togo.User) (bool, error) {
	return ur.IsExceedPerDayFn(ctx, u)
}
