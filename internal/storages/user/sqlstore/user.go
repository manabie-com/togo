package sqlstore

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages/user/model"
)

type UserStore struct {
	*sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

// FindByID find user by id
func (s *UserStore) FindByID(ctx context.Context, userID sql.NullString) (*model.User, error) {
	stmt := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := s.QueryRowContext(ctx, stmt, userID)

	user := new(model.User)
	err := row.Scan(&user.ID, &user.Password, &user.MaxTodo)
	if err != nil {
		return nil, err
	}
	
	return user, err
}
