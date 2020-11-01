package sqlite

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entities"
)

// LiteDB for working with sqlite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]entities.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task
	for rows.Next() {
		t := entities.Task{}
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
func (l LiteDB) AddTask(ctx context.Context, t entities.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &entities.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

// GetMaxTaskTodo get the number of limit task accordinate with userID
func (l LiteDB) GetMaxTaskTodo(ctx context.Context, userID string) (int, error) {
	stmt := `SELECT max_todo FROM "users" WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	var maxTask int
	err := row.Scan(&maxTask)
	return maxTask, err
}
