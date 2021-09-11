package sqlite

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

const (
	sqlValidateUser = `SELECT id FROM users WHERE id = ? AND password = ?`
	sqlRetrieveTasks = `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	sqlAddTask = `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
)

// liteDB for working with sqllite
type liteDB struct {
	db *sql.DB
}

func NewLiteDB(db *sql.DB) *liteDB {
	return &liteDB{
		db: db,
	}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *liteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	rows, err := l.db.QueryContext(ctx, sqlRetrieveTasks, userID, createdDate)
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

// AddTask adds a new task to db
func (l *liteDB) AddTask(ctx context.Context, t *storages.Task) error {
	_, err := l.db.ExecContext(ctx, sqlAddTask, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *liteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	row := l.db.QueryRowContext(ctx, sqlValidateUser, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
