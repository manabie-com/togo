package domain

import (
	"context"
	"github.com/manabie-com/togo/internal/domain/entity"
)

type UserRepository interface {
	// GetUser by username
	GetUser(ctx context.Context, username string) (*entity.User, error)
}
