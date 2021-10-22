package pg

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/manabie-com/togo/internal/storages"
)

// PgDB for working with postgres
type PgDB struct {
	DB  *pgxpool.Pool
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PgDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]storages.Task, error) {
	var tasks []storages.Task
	query := "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2"
	err := pgxscan.Select(ctx, l.DB, &tasks, query, userID, createdDate)
	return tasks, err
}

// AddTask adds a new task to DB
func (l *PgDB) AddTask(ctx context.Context, t *storages.Task) error {
	query := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.Exec(ctx, query, t.ID, t.Content, t.UserID, t.CreatedDate)
	return err
}

// ValidateUser returns tasks if match userID AND password
func (l *PgDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	u := storages.User{}
	query := "SELECT id FROM users WHERE id = $1 AND password = $2"
	err := pgxscan.Get(ctx, l.DB, &u, query, userID, pwd)
	if err != nil {
		return false
	}

	return true
}
