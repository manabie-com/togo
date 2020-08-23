package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/phuwn/togo/internal/storages"
	"golang.org/x/crypto/bcrypt"
)

// Storage object for working with database
type Storage struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (s *Storage) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := s.DB.QueryContext(ctx, stmt, userID, createdDate)
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
func (s *Storage) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (s *Storage) ValidateUser(ctx context.Context, userID sql.NullString, pwd string) bool {
	stmt := `SELECT id, password FROM users WHERE id = $1`
	row := s.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.Password)
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return false
	}

	return true
}
