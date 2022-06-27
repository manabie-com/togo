package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"manabie/togo/internal/model"
)

const (
	CreateUserStmt = `INSERT INTO users (id, password) VALUES ($1, $2)`
	FindUserStmt   = `SELECT id, password, max_todo FROM users WHERE id = $1`
)

type userStore struct {
	DB *sql.DB
}

func (s *userStore) Create(ctx context.Context, u *model.User) error {
	if _, err := s.DB.ExecContext(ctx, CreateUserStmt, &u.ID, &u.Password); err != nil {
		return err
	}

	return nil
}
func (s *userStore) FindUser(ctx context.Context, userID string) (*model.User, error) {
	row := s.DB.QueryRowContext(ctx, FindUserStmt, userID)
	u := &model.User{}
	if err := row.Scan(&u.ID, &u.Password, &u.MaxTodo); err != nil {
		return nil, err
	}

	return u, nil
}

func NewUserStore(db *sql.DB) model.UserStore {
	return &userStore{
		DB: db,
	}
}
