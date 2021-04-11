package postgres

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"

	"github.com/jackc/pgx/v4"
)

// PostgresDB for working with PostgreSQL
type PostgresDB struct {
	DB *pgx.Conn
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (pg *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	return nil, nil
}

// AddTask adds a new task to DB
func (pg *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (pg *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	return false
}
