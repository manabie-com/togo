package sqllite

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
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
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	var max_todo int
	var current_count int
	// Get max_todo
	max_todo_stmt := `SELECT max_todo FROM users WHERE id = ?`
	max_todo_row := l.DB.QueryRowContext(ctx, max_todo_stmt, &t.UserID)
	err := max_todo_row.Scan(&max_todo)
	if err != nil {
		return err
	}
	//Get current count for this created date
	stmt := `SELECT count(created_date) FROM Tasks WHERE created_date = ?`
	row := l.DB.QueryRowContext(ctx, stmt, &t.CreatedDate)
	err = row.Scan(&current_count)
	if err != nil {
		log.Fatal(err)
	}

	if max_todo <= current_count {
		return fmt.Errorf("Exceed")
	}

	//Every is good and a new record will be inserted into database
	stmt = `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err = l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
