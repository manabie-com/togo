package sqllite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
// createDate is optional
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT t.id, t.content, t.user_id, t.status_code, st.name, t.created_date, t.updated_at
			FROM tasks t 
				JOIN task_status st ON st.code = t.status_code
			WHERE t.user_id = ?`
	if len(createdDate.String) > 0 {
		stmt += "AND t.created_date = \"" + createdDate.String + "\""
	}
	rows, err := l.DB.QueryContext(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*storages.Task
	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.StatusCode, &t.StatusName, &t.CreatedDate, &t.UpdatedAt)
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
	stmt := `INSERT INTO tasks (id, content, user_id, status_code, created_date, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, t.ID, &t.Content, t.UserID, t.StatusCode, t.CreatedDate, t.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTask update existing task
func (l *LiteDB) UpdateTask(ctx context.Context, t *storages.Task) error {
	stmt := `UPDATE tasks SET `

	hasContent := false
	hasStatus := false
	if len(t.Content) > 0 {
		stmt += "content=\"" + t.Content + "\""
		hasContent = true
	}
	if len(t.StatusCode) > 0 {
		if hasContent {
			stmt += ","
		}
		stmt += "status_code=\"" + t.StatusCode + "\""
		hasStatus = true
	}

	if hasContent || hasStatus {
		stmt += ","
	}

	stmt += "updated_at=? WHERE id = ?"

	fmt.Println(stmt)

	_, err := l.DB.ExecContext(ctx, stmt, t.UpdatedAt, t.ID)
	if err != nil {
		return err
	}

	return nil
}

// Get user max todo
func (l *LiteDB) GetUserMaxTodo(ctx context.Context, userID sql.NullString) (int, error) {
	stmt := `SELECT max_todo FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID)

	u := &storages.User{}

	err := row.Scan(&u.MaxTodo)
	if err != nil {
		return 0, err
	}

	return u.MaxTodo, nil
}

// Count user today task adds a new task to DB
func (l *LiteDB) CountTodayTask(ctx context.Context, userID sql.NullString) (int, error) {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date = ?`
	currentTime := time.Now()
	rows, err := l.DB.QueryContext(ctx, stmt, userID, currentTime.Format("2006-01-02"))
	if err != nil {
		fmt.Println("Count error ", err)
		return 0, err
	}

	defer rows.Close()

	var count int

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, err
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
