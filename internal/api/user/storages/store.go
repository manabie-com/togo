package storages

import (
	"context"
)

//go:generate mockgen -package mock -destination mock/store_mock.go . Store
type Store interface {
	Get(ctx context.Context, userID string) (*User, error)
}
