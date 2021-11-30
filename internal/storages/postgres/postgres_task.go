package postgres

import (
	"context"
	"database/sql"
	"fmt"
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
	CountTaskPerDay(ctx context.Context, userID string, createdDate time.Time) (uint8, error)
}

func NewPostgresDB(Conn *sql.DB) TaskDB {
	return &postgresDB{Conn}
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *postgresDB) RetrieveTasks(ctx context.Context, userID string, createdDate time.Time) (result []storages.Task, err error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND to_char(created_date,'YYYY-MM-DD') = $2`
	rows, err := p.Conn.QueryContext(ctx, stmt, userID, createdDate.Format("2006-01-02"))
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

	stmt := `INSERT INTO tasks (content, user_id, created_date) VALUES ($1, $2, $3) RETURNING id`
	rows := p.Conn.QueryRowContext(ctx, stmt, &t.Content, &t.UserID, &t.CreatedDate)
	var id int64
	err := rows.Scan(&id)
	if err != nil {
		return err
	}
	t.ID = id
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (p *postgresDB) ValidateUser(ctx context.Context, userID string, pwd string) (bool, error) {
	fmt.Println(userID + ":" + pwd)
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.Conn.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
func (p *postgresDB) CountTaskPerDay(ctx context.Context, userID string, createdDate time.Time) (uint8, error) {
	stmt := `SELECT count(id) FROM tasks WHERE user_id = $1 AND to_char(created_date,'YYYY-MM-DD') = $2`
	row := p.Conn.QueryRowContext(ctx, stmt, userID, createdDate.Format("2006-01-02"))
	var total uint8
	err := row.Scan(&total)
	return total, err
}
