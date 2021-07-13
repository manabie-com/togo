package storages

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entity"
)

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *userStorage {
	return &userStorage{
		db: db,
	}
}

// Find returns user if found.
func (l *userStorage) FindByID(ctx context.Context, userID string) (*entity.User, error) {
	stmt := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := l.db.QueryRowContext(ctx, stmt, userID)
	u := &entity.User{}
	err := row.Scan(&u.ID, &u.Password, &u.MaxTodoPerday)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, err
}
