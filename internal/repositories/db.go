package repositories

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/models"
)

// LiteDB for working with sqllite
type DB struct {
	DB *sql.DB
}

// Returns the max_todo of the user if match userID
func (l *DB) RetrieveMaxToDoSetting(ctx context.Context, userID sql.NullString) int {
	query := `SELECT * FROM users WHERE id = $1 LIMIT 1`
	rows, err := l.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return 0
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		t := &models.User{}
		err := rows.Scan(&t.ID, &t.Password, &t.MaxToDo)
		if err != nil {
			return 0
		}
		users = append(users, t)
	}

	if err := rows.Err(); err != nil {
		return 0
	}

	if len(users) == 0 {
		return 0
	}

	return users[0].MaxToDo
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *DB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*models.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
	rows, err := l.DB.QueryContext(ctx, stmt, userID, createdDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		t := &models.Task{}
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
func (l *DB) AddTask(ctx context.Context, t *models.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *DB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &models.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
