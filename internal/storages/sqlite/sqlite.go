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
	stmt, err := l.DB.PrepareContext(ctx, `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, userID, createdDate)
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
	stmt, err := l.DB.PrepareContext(ctx, `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, t.ID, t.Content, t.UserID, t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt, err := l.DB.PrepareContext(ctx, `SELECT id FROM users WHERE id = ? AND password = ?`)
	if err != nil {
		return false
	}
	row := stmt.QueryRowContext(ctx, userID, pwd)
	u := &entities.User{}
	err = row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

// GetMaxTaskTodo get the number of limit task accordinate with userID
func (l LiteDB) GetMaxTaskTodo(ctx context.Context, userID string) (int, error) {
	var maxTask int
	stmt, err := l.DB.PrepareContext(ctx, `SELECT max_todo FROM "users" WHERE id = ?`)
	if err != nil {
		return maxTask, err
	}
	row := stmt.QueryRowContext(ctx, userID)
	err = row.Scan(&maxTask)
	return maxTask, err
}
