package sqllite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/manabie-com/togo/internal/storages"
)

// PostgresDB for working with postgres
type PostgresDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (p *PostgresDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
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

// CountTasks returns number of tasks if match userID AND createDate.
func (p *PostgresDB) CountTasks(ctx context.Context, userID, createdDate sql.NullString) (uint, error) {
	stmt := `SELECT count(id) FROM tasks WHERE user_id = $1 AND created_date = $2`
	row, err := p.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var numOfTasks uint
	row.Next()
	err = row.Scan(&numOfTasks)
	if err != nil {
		return 0, err
	}

	if err := row.Err(); err != nil {
		return 0, err
	}

	return numOfTasks, nil
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
func (p *PostgresDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := p.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// GetMaxToDo returns max to do task per day if match userID
func (p *PostgresDB) GetMaxToDo(ctx context.Context, userID sql.NullString) (uint, error) {
	stmt := `SELECT max_todo FROM users WHERE id = $1`
	row := p.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.MaxTodo)
	if err != nil {
		return 0, err
	}

	return u.MaxTodo, nil
}
