package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
)

// PG for working with sqlpg
type PG struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PG) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := l.DB.Query(stmt, userID, createdDate)
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
func (l *PG) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.Exec(stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

//
func (l *PG) CheckAddTask(ctx context.Context, userID string, maxTodo int, createDate string) bool {
	stmt := `SELECT count(id) FROM tasks WHERE user_id = $1 AND created_date = $2`
	var count int
	row := l.DB.QueryRow(stmt, userID, createDate)
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	if count < maxTodo {
		return false
	}
	return true
}

func (l *PG) GetMaxTodoUser(ctx context.Context, userID string) int {
	stmt := `SELECT max_todo FROM users WHERE id = $1`
	row := l.DB.QueryRow(stmt, userID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

// ValidateUser returns tasks if match userID AND password
func (l *PG) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.DB.QueryRow(stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
