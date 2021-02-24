package sqlstore

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages/model"
)

// ValidateUser returns tasks if match userID AND password
func (s *Store) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := s.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &model.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
