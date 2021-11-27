package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/opentracing/opentracing-go/log"
)

type postgresDB struct {
	Conn *sql.DB
}
type TaskDB interface {
	RetrieveTasks(ctx context.Context, userID string, createdDate time.Time) ([]storages.Task, error)
	AddTask(ctx context.Context, t *storages.Task) error
	ValidateUser(ctx context.Context, userID string, pwd string) (bool, error)
}

func NewPostgresDB(Conn *sql.DB) TaskDB {
	return &postgresDB{Conn}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *postgresDB) RetrieveTasks(ctx context.Context, userID string, createdDate time.Time) (result []storages.Task, err error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := p.Conn.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()
	result = make([]storages.Task, 0)
	for rows.Next() {
		t := storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

// AddTask adds a new task to DB
func (p *postgresDB) AddTask(ctx context.Context, t *storages.Task) error {

	stmt := `INSERT INTO tasks (content, user_id, created_date) VALUES (?, ?, ?)`
	_, err := p.Conn.ExecContext(ctx, stmt, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (p *postgresDB) ValidateUser(ctx context.Context, userID string, pwd string) (bool, error) {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := p.Conn.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}
