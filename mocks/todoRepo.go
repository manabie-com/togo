package mocks

import (
	"context"

	"github.com/lawtrann/togo"
)

type TodoRepo struct {
	AddFn            func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error)
	AddWithNewUserFn func(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error)
}

func (tp *TodoRepo) Add(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
	return tp.AddFn(ctx, t, u)
}

func (tp *TodoRepo) AddWithNewUser(ctx context.Context, t *togo.Todo, u *togo.User) (*togo.Todo, error) {
	return tp.AddWithNewUserFn(ctx, t, u)
}
