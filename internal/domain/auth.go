package domain

import (
	"context"
)

// AuthUseCase auth use case
type AuthUseCase interface {
	// SignIn sign in with username and password
	SignIn(ctx context.Context, username string, password string) (string, error)
}
