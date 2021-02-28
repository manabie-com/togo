package sqlstore

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages/model"
)

// FindByID find user by id
func (s *Store) FindByID(ctx context.Context, userID sql.NullString) (*model.User, error) {
	stmt := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := s.DB.QueryRowContext(ctx, stmt, userID)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Password, &user.MaxTodo)
	
	return user, err
}
