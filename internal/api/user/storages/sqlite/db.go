package sqllite

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/api/user/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) Get(ctx context.Context, userID string) (*storages.User, error) {
	id := sql.NullString{
		String: userID,
		Valid:  true,
	}
	query := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := l.DB.QueryRowContext(ctx, query, id)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.Password, &u.MaxTodo)
	if err != nil {
		return u, err
	}
	return u, nil
}
