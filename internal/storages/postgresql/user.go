package postgresql

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
)

// PSQLUserRespsitory ...
type PSQLUserRespsitory struct {
	db *sql.DB
}

type psqlUser struct {
	UserID   string
	Password string
}

// NewPSQLUserRepository ...
func NewPSQLUserRepository(db *sql.DB) *PSQLUserRespsitory {
	return &PSQLUserRespsitory{
		db: db,
	}
}

// GetUserByUserID ...
func (psql *PSQLUserRespsitory) GetUserByUserID(ctx context.Context, userID string) (*entities.User, error) {
	var foundUser psqlUser
	err := psql.db.QueryRow("SELECT id, password from users WHERE id = $1;", userID).Scan(&foundUser.UserID, &foundUser.Password)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &entities.User{
			ID:       foundUser.UserID,
			Password: foundUser.Password,
		}, nil
	default:
		return nil, storages.ErrInternalError
	}
}
