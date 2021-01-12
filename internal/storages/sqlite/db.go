package sqllite

import (
	"github.com/manabie-com/togo/internal/storages"

	"context"
	"database/sql"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
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

// AddTask adds a new task to DB
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)

	if err != nil {
		return false
	}

	return true
}

// GetUser get a user by id
func (l *LiteDB) GetUser(ctx context.Context, userID sql.NullString) (*storages.User, error) {
	stmt := `SELECT id, password, max_todo FROM users WHERE id = $1`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.Password, &u.MaxTodo)

	if err != nil {
		return nil, err
	}

	return u, nil
}

// CountTasks count tasks that match userID AND created date
func (l *LiteDB) CountTasks(ctx context.Context, userID, createdDate sql.NullString) (int, error) {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, createdDate)
	total := 0
	err := row.Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
