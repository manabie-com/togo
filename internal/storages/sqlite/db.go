package sqllite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/banhquocdanh/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date = ?`
	userIDNotNullString := sql.NullString{
		String: userID,
		Valid:  true,
	}
	createdDateNotNullString := sql.NullString{
		String: createdDate,
		Valid:  true,
	}

	rows, err := l.DB.QueryContext(ctx, stmt, userIDNotNullString, createdDateNotNullString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}
	var tasks []*storages.Task

	for rows.Next() {
		t := &storages.Task{}
		err := rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// AddTask adds a new task to DB
func (l *LiteDB) addTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func (l *LiteDB) countTaskPerDayByUserID(ctx context.Context, userID string, createdDate string) (uint, error) {
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date = ?`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count uint
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Printf("Err: %s\n", err)
			return count, err
		}
	}

	return count, nil
}

// AddTask adds a new task to DB
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) error {
	user, err := l.getUser(ctx, t.UserID)
	if err != nil {
		return err
	}
	count, err := l.countTaskPerDayByUserID(ctx, user.ID, t.CreatedDate)
	if err != nil {
		return err
	}

	if user.MaxTodo <= count {
		return fmt.Errorf("max todo task")
	}

	return l.addTask(ctx, t)
}

// getUser returns tasks if match with userID
func (l *LiteDB) getUser(ctx context.Context, userID string) (*storages.User, error) {
	userIDNotNullString := sql.NullString{String: userID, Valid: true}
	stmt := `SELECT * FROM users WHERE id = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userIDNotNullString)
	u := &storages.User{}
	err := row.Scan(&u.ID, &u.Password, &u.MaxTodo)

	return u, err
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID, pwd string) bool {
	userIDNotNullString := sql.NullString{String: userID, Valid: true}
	pwdNotNullString := sql.NullString{String: pwd, Valid: true}
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userIDNotNullString, pwdNotNullString)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
