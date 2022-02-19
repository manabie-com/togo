package repository

import "context"

// TaskLimitRepository repository interface
type TaskLimitRepository interface {
	Increase(ctx context.Context, userID uint, limit int) (int, error)
}
