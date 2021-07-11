package database

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

// Find returns user if found.
func (l *userRepository) Find(ctx context.Context, userID, pwd string) (*entity.User, error) {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &entity.User{}
	err := row.Scan(&u.ID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, nil
}
