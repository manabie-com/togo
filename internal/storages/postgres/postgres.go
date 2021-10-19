package postgres

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

// PostgresDB for working with sqllite
type postgresDB struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) storages.Repository {
	return &postgresDB{db}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *postgresDB) RetrieveTasks(ctx context.Context, userID uint, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := l.db.QueryContext(ctx, stmt, userID, createdDate)
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
func (l *postgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.db.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByUsername returns user if match username
func (l *postgresDB) GetUserByUsername(ctx context.Context, username sql.NullString) (*storages.User, error) {
	stmt := "SELECT id, username, max_todo FROM users WHERE username = $1"
	row := l.db.QueryRowContext(ctx, stmt, username)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.Username, &u.MaxTodo)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// ValidateUser returns tasks if match username AND password
func (l *postgresDB) ValidateUser(ctx context.Context, username, pwd sql.NullString) bool {
	stmt := "SELECT id FROM users WHERE username = $1 AND password = $2"
	row := l.db.QueryRowContext(ctx, stmt, username, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	return err == nil
}
