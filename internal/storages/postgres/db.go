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
	const query = "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2"

	rows, err := pg.DB.Query(ctx, query, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		var t storages.Task

		if err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (pg *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (pg *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	const query = "SELECT id FROM users WHERE id = $1 AND password = $2"

	var u storages.User

	if err := pg.DB.QueryRow(ctx, query, userID, pwd).Scan(&u.ID); err != nil {
		return false
	}

	return true
}
