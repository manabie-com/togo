package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/api/user/storages"
)

// PostgresDB for working with postgres
type PostgresDB struct {
	DB *sql.DB
}

// ValidateUser returns tasks if match userID AND password
func (l *PostgresDB) Get(ctx context.Context, userID string) (*storages.User, error) {
	id := sql.NullString{
		String: userID,
		Valid:  true,
	}
	query := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := l.DB.QueryRowContext(ctx, query, id)
	user := &storages.User{}
	err := row.Scan(&user.ID, &user.Password, &user.MaxTodo)
	if err != nil {
		return user, err
	}
	return user, nil
}
