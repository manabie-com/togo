package sqllite

import (
	"context"
	"database/sql"
	"log"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date >= ?`
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
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// GetUserProfile gets user profile
func (l *LiteDB) GetUserProfile(ctx context.Context, userID sql.NullString) (*storages.User, error) {
	log.Println("UserID", userID)
	stmt := `SELECT id, max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.MaxTodo)

	return u, err
}

// CountUserTasks count  user tasks by date
func (l *LiteDB) CountUserTasksByDate(ctx context.Context, userID, created_date sql.NullString) (int, error) {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date >= ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, created_date)
	var count int
	err := row.Scan(&count)

	return count, err
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	// log.Print(u)
	if err != nil {
		return false
	}

	return true
}
