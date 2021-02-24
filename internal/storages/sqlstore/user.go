package sqlstore

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages/model"
	"github.com/manabie-com/togo/pkg/common/crypto"
)

// ValidateUser returns tasks if match userID AND password
func (s *Store) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id, password FROM users WHERE id = $1`
	row := s.DB.QueryRowContext(ctx, stmt, userID)
	u := &model.User{}

	err := row.Scan(&u.ID, &u.Password)
	if err != nil {
		return false
	}

	return crypto.CheckPasswordHash(pwd.String, u.Password)
}
