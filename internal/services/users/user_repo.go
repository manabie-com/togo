package users

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	OldPasswords []string  `json:"old_passwords"`
}

type UserRepo interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
}
