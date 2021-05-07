package storages

import (
	"context"
	"database/sql"
	"github.com/google/martian/log"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

func (l *LiteDB) CountTasks(ctx context.Context, userID, date string) (int32, error) {
	var (
		numOfTask sql.NullInt32
	)
	stmt := `SELECT COUNT(t.id) FROM tasks t WHERE t.id = ? AND t.created_date = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, date)
	if row.Err() != nil {
		return 0, row.Err()
	}
	err := row.Scan(&numOfTask)
	return numOfTask.Int32, err
}

func (l *LiteDB) GetMaxTodo(ctx context.Context, userID string) (int32, error) {
	var maxTodo sql.NullInt32
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)
	if row.Err() != nil {
		return 0, row.Err()
	}
	err := row.Scan(&maxTodo)
	return maxTodo.Int32, err
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		t := &Task{}
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
func (l *LiteDB) AddTask(ctx context.Context, t *Task, callback func(string, string) error) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	// new transaction
	tx, err := l.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate); err != nil {
		tx.Rollback()
		return err
	}
	// call callback function if err occurs then rollback
	if err = callback(t.UserID, t.CreatedDate); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

func (l *LiteDB) Close() {
	if err := l.DB.Close(); err != nil {
		log.Errorf("error while closing db - %s", err.Error())
	}
}