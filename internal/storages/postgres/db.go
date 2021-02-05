package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils"
)

// Postgres for working with Postgresql
type Postgres struct {
	DB *sql.DB
}

//func (p *Postgres) GetDB() *Postgres {
//	once.Do(func() {
//		println("Init db here")
//		err := p.connect()
//		if err != nil {
//			utils.Error(err.Error())
//		}
//	})
//	println("Return db here")
//	return _db
//}
//
//func (p *Postgres) connect() error {
//	var err error
//	pgInfo := config.GetConfig().GetConnString()
//	println("connect db here")
//	p.DB, err = sql.Open("postgres", pgInfo)
//	if err != nil {
//		return err
//	}
//
//	err = p.DB.Ping()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *Postgres) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := p.DB.QueryContext(ctx, stmt, userID, createdDate)
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
func (p *Postgres) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := p.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (p *Postgres) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		utils.Error(err.Error())
		return false
	}

	return true
}

// GetUserInfo returns user info when existing userID or null if vice versa
func (p *Postgres) GetUserInfo(ctx context.Context, userID sql.NullString) *storages.User {
	stmt := `SELECT id, max_todo FROM users WHERE id = $1`
	row := p.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.MaxTodo)
	if err != nil {
		return nil
	}

	return u
}

// CountTasks returns the number of tasks that matched userID AND createDate.
func (p *Postgres) CountTasks(ctx context.Context, userID, createdDate sql.NullString) (int, error) {
	stmt := `SELECT COUNT(id) as total FROM tasks WHERE user_id = $1 AND created_date = $2`
	row := p.DB.QueryRowContext(ctx, stmt, userID, createdDate)

	var total int
	err := row.Scan(&total)
	if err != nil {
		return -1, err
	}

	return total, nil
}
