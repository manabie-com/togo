package storages

import (
	"context"
)

type Store interface {
	ValidateUser(ctx context.Context, userID string) (*User, error)
}
