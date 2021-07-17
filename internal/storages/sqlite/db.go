package sqllite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/storages"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *LiteDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ?`

	// only filter by created date if exists in the request
	if createdDate.String != "" {
		stmt += " AND created_date = ?"
	}
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
	// check allowance
	totalAddedToday, _ := l.GetTotalAddedTasksOfUserByDate(ctx, t.UserID, t.CreatedDate)
	userAllowance, _ := l.GetUserMaxToDo(ctx, t.UserID)
	if totalAddedToday == userAllowance {
		return errors.New("You already reached the maximum tasks that you can add for a day. Please come back tomorrow.")
	}

	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *LiteDB) ValidateUser(ctx context.Context, userID sql.NullString, pwd sql.NullString) bool {
	stmt := `SELECT id, password FROM users WHERE id = ? LIMIT 1`

	row := l.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}

	err := row.Scan(&u.ID, &u.Password)
	if err != nil {
		return false
	}

	return helpers.CheckPasswordHash(pwd.String, u.Password)
}

func (l *LiteDB) GetUserMaxToDo(ctx context.Context, userID string) (int, error) {
	stmt := `SELECT max_todo FROM users WHERE id = ? LIMIT 1`

	row := l.DB.QueryRowContext(ctx, stmt, userID)
	u := &storages.User{}

	err := row.Scan(&u.MaxToDo)
	if err != nil {
		return 0, err
	}

	return u.MaxToDo, nil
}

func (l *LiteDB) GetTotalAddedTasksOfUserByDate(ctx context.Context, userID, createdDate string) (int, error) {
	stmt := `SELECT COUNT(id) FROM tasks WHERE user_id = ? AND created_date = ?`

	row := l.DB.QueryRow(stmt, userID, createdDate)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
