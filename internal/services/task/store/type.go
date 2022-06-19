package store

import (
	"github.com/google/uuid"
	"time"
	"togo/internal/services/task/domain"
)

type UserRepository interface {
	Save(entity *domain.User) error
	FindByID(uuid uuid.UUID) (*domain.User, error)
}

type TaskRepository interface {
	Save(entity *domain.Task) error
	Count(req CountTasksRequest) (int, error)
}

type CountTasksRequest struct {
	UserID *uuid.UUID
	Day    *time.Time
}
