package mocks

import (
	"context"

	"github.com/lawtrann/togo"
)

type TodoService struct {
	AddFn func(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error)
}

func (ts *TodoService) Add(ctx context.Context, t *togo.Todo, username string) (*togo.Todo, error) {
	return ts.AddFn(ctx, t, username)
}
