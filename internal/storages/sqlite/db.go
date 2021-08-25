package sqllite

import (
	"context"
	"database/sql"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type (
	LiteDB interface {
		RetrieveTasks(ctx context.Context, userID, createdDate string) ([]storages.Task, error)
		AddTask(ctx context.Context, task storages.Task) error
		ValidateUser(ctx context.Context, userID, pwd string) bool
		GetMaxTodo(ctx context.Context, userID string) (uint32, error)
	}
	liteDB struct {
		db *sql.DB
	}
)

func NewLiteDB(db *sql.DB) LiteDB {
	return &liteDB{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *liteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]storages.Task, error) {
	ctx, canncel := context.WithTimeout(ctx, 60*time.Second)
	defer canncel()

	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.db.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []storages.Task
	for rows.Next() {
		task := storages.Task{}
		err := rows.Scan(&task.ID, &task.Content, &task.UserID, &task.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *liteDB) AddTask(ctx context.Context, task storages.Task) error {
	ctx, canncel := context.WithTimeout(ctx, 60*time.Second)
	defer canncel()

	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.db.ExecContext(ctx, stmt, task.ID, task.Content, task.UserID, task.CreatedDate)
	if err != nil {
		return err
	}
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *liteDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	ctx, canncel := context.WithTimeout(ctx, 60*time.Second)
	defer canncel()

	var id string
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	err := l.db.QueryRowContext(ctx, stmt, userID, pwd).Scan(&id)
	return err == nil
}

func (l *liteDB) GetMaxTodo(ctx context.Context, userID string) (uint32, error) {
	ctx, canncel := context.WithTimeout(ctx, 60*time.Second)
	defer canncel()

	var maxTodo uint32
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	err := l.db.QueryRowContext(ctx, stmt, userID).Scan(&maxTodo)
	if err != nil {
		return 0, err
	}
	return maxTodo, nil
}
