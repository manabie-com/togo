package storages

import (
	"context"
	"errors"

	"github.com/manabie-com/togo/internal/entities"
)

// Implement repository pattern
// which decouple the storage layer from the business logic layer

// TaskRepository ...
type TaskRepository interface {
	SaveTask(ctx context.Context, task entities.Task) (*entities.Task, error)
	GetTasksByUserIDAndDate(ctx context.Context, userID, createdDate string) ([]*entities.Task, error)
	CountTasksOfUserByDate(ctx context.Context, userID, createdDate string) (int, error)
}

// UserRepository ...
type UserRepository interface {
	GetUserByUserID(ctx context.Context, userID string) (*entities.User, error)
	GetUserTaskLimit(ctx context.Context, userID string) (int, error)
}

// Common error
var (
	ErrUserNotFound  = errors.New("User not found")
	ErrInternalError = errors.New("Cannot perform action against database")
)
