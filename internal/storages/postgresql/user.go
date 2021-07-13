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
func (l *userRepository) FindByID(ctx context.Context, userID string) (*entity.User, error) {
	stmt := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := l.db.QueryRowContext(ctx, stmt, userID)
	u := &entity.User{}
	err := row.Scan(&u.ID, &u.Password, &u.MaxTodoPerday)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, err
}
