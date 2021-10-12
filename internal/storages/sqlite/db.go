package sqllite

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveUserTaskLimit returns max number of tasks that can be created if match ID.
func (l *LiteDB) RetrieveUserTaskLimit(ctx context.Context, userID sql.NullString) (int, error) {
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)

	var taskCount int
	err := row.Scan(&taskCount)

	return taskCount, err
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// CountTasks returns total number of tasks if match userID AND createDate.
func (l *LiteDB) CountTasks(ctx context.Context, userID, createdDate sql.NullString) (int, error) {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, createdDate)

	var taskCount int
	err := row.Scan(&taskCount)

	return taskCount, err
}

// AddTask adds a new task to DB
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
