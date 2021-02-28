package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

type PostgresDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *PostgresDB) RetrieveTasks(ctx context.Context, userID sql.NullString, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := p.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}

		parsedTime, _ := time.Parse(time.RFC3339, t.CreatedDate)
		t.CreatedDate = parsedTime.Format("2006-01-02")
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (p *PostgresDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := p.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (p *PostgresDB) ValidateUser(ctx context.Context, userID sql.NullString, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
