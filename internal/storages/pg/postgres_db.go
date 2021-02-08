package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"
)

type PostgresDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, *config.ErrorInfo) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := p.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, &config.ErrorInfo{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, &config.ErrorInfo{
				Err:        err,
				StatusCode: http.StatusInternalServerError,
			}
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, &config.ErrorInfo{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (p *PostgresDB) AddTask(ctx context.Context, t *storages.Task) *config.ErrorInfo {

	tx, err := p.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return &config.ErrorInfo{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	stmt := `SELECT max_todo FROM users WHERE id = $1`
	rows, err := tx.QueryContext(ctx, stmt, t.UserID)
	if err != nil {
		tx.Rollback()
		return &config.ErrorInfo{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	var maxTodo int
	existUser := false
	for rows.Next() {
		existUser = true
		err = rows.Scan(&maxTodo)
		if err != nil {
			tx.Rollback()
			return &config.ErrorInfo{
				Err:        err,
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	if !existUser {
		tx.Rollback()
		return &config.ErrorInfo{
			Err:        errors.New(fmt.Sprintf("Not found any user with id %s", t.UserID)),
			StatusCode: http.StatusNotFound,
		}
	}

	stmt = `SELECT count(id) FROM tasks where user_id = $1 and created_date = $2`
	rows, err = tx.QueryContext(ctx, stmt, t.UserID, t.CreatedDate)
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			tx.Rollback()
			return &config.ErrorInfo{
				Err:        err,
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	if count >= maxTodo {
		tx.Rollback()
		return &config.ErrorInfo{
			Err:        errors.New(fmt.Sprintf("User with id %s has enough %d task per day", t.UserID, maxTodo)),
			StatusCode: http.StatusBadRequest,
		}
	}

	stmt = `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		tx.Rollback()
		return &config.ErrorInfo{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	err = tx.Commit()
	if err != nil {
		return &config.ErrorInfo{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}
	return nil
}

// ValidateUser returns tasks if match userID AND password
func (p *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.DB.QueryRowContext(ctx, stmt, userID.String, pwd.String)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
