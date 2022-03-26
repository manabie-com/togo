package repository

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/models"
)

// DatabaseRepository interface is the definition for methods collection
// that interact directly with the database
type DatabaseRepository interface {
	ValidateUser(ctx context.Context, username, password sql.NullString) bool
	GetUserByUserName(ctx context.Context, username sql.NullString) (*models.User, error)
	AddTask(ctx context.Context, task *models.Task) error
	RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*models.Task, error)
}
