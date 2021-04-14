package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserConfig struct {
	ID        int64     `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	TaskLimit int       `json:"task_limit"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserConfigRepo interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*UserConfig, error)
}
