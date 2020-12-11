package sqllite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/manabie-com/togo/internal/storages"
)

// Databaser is responsible for talking with database
type Databaser interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
	AddTask(ctx context.Context, t *storages.Task) (int64, error)
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
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

// AddTask adds a new task to DB if the user's daily-limit has not been reached. It returns
// the number of tasks added (0 or 1) and the error if exists. The isolation level of the underlying
// transaction is set to SERIALIZABLE to ensure concurrency correctness.
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) (int64, error) {
	tx, err := l.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return 0, err
	}
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) ` +
		`SELECT ?, ?, ?, ? ` +
		`WHERE (SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date = ?) ` +
		`< (SELECT max_todo FROM users WHERE id = ?)`
	result, err := tx.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate, &t.UserID, &t.CreatedDate, &t.UserID)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			fmt.Printf("Update failed: %v, unable to rollback: %v\n", err, rbErr)
			return 0, rbErr
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return result.RowsAffected()
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
