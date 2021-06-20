package sqllite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {

	stmt := `SELECT * FROM tasks WHERE user_id = ? AND created_date = ?`
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
	stmt := `INSERT INTO tasks (id, content, user_id, created_date, updated_date, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate, "", &t.UserID, "")
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

// Add new user
func (l *LiteDB) RegisterUser(ctx context.Context, userId string, pwd string) error {
	stmt := `INSERT INTO users (id, password) VALUES (?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, userId, pwd)
	if err != nil {
		return err
	}

	return nil
}

func (l *LiteDB) GetUsers(ctx context.Context) ([]*storages.User, error) {
	stmt := `select id from users`
	rows, err := l.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*storages.User
	for rows.Next() {
		u := &storages.User{}
		err := rows.Scan(&u.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (l *LiteDB) GetListTasks(ctx context.Context) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks`
	rows, err := l.DB.QueryContext(ctx, stmt)
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

func (l *LiteDB) RetrieveTasks1(ctx context.Context, userID string, createdDate string) ([]*storages.Task, error) {
	stmt := `SELECT * FROM tasks WHERE user_id = ? AND created_date = ?`
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

func (l *LiteDB) UpdateTask(ctx context.Context, t *storages.Task) error {
	stmt := `UPDATE tasks set content = ? where id = ?`
	fmt.Println("content", &t.Content)
	_, err := l.DB.ExecContext(ctx, stmt, &t.Content, &t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (l *LiteDB) GetTaskById(ctx context.Context, taskId string) ([]*storages.Task, error) {
	stmt := `select * from tasks where id = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, taskId)
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

func (l *LiteDB) DeleteTask(ctx context.Context, taskId string) error {
	stmt := `DELETE FROM tasks WHERE id = ?`
	_, err := l.DB.ExecContext(ctx, stmt, taskId)
	if err != nil {
		return err
	}

	return nil
}
