package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/storages"
)

// Vendor for working with sqllite & postgre
type Vendor struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (vendor *Vendor) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date, status  FROM tasks WHERE user_id = $1 AND created_date = $2 AND status <> $3`
	rows, err := vendor.DB.QueryContext(ctx, stmt, userID, createdDate, storages.DELETED)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate, &t.Status)
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
func (vendor *Vendor) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date,status) VALUES ($1, $2, $3, $4, $5)`
	_, err := vendor.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate, storages.ACTIVE)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (vendor *Vendor) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := vendor.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}
	return true
}

func (vendor *Vendor) GetUserById(ctx context.Context, userID string) *storages.User {
	stmt := `SELECT id, max_todo FROM users WHERE id = $1`
	row := vendor.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.MaxTodo)
	if err != nil {
		return nil
	}
	return u
}


func (vendor *Vendor) UpdateTodoStatus(ctx context.Context, taskId sql.NullString, status storages.TaskType) error {
	stmt := `UPDATE tasks SET status = $1 WHERE id = $2`
	_, err := vendor.DB.ExecContext(ctx, stmt, status ,taskId)
	if err != nil {
		return err
	}
	return nil
}
