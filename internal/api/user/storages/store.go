package storages

import (
	"context"
)

type Store interface {
	Get(ctx context.Context, userID string) (*User, error)
}
