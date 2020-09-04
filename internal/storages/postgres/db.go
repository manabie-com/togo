package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

// PqDB ...
type PqDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PqDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2`
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
func (l *PqDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *PqDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}

// GetUserMaxTask get max task that user can add per day
func (l *PqDB) GetUserMaxTask(ctx context.Context, userID string) (int, error) {
	stmt := `SELECT max_todo FROM users WHERE id = $1`
	row := l.DB.QueryRowContext(ctx, stmt, userID)

	var maxTask int
	err := row.Scan(&maxTask)

	if err != nil {
		return 0, err
	}

	return maxTask, nil
}

// GetUserTodayTask get number of task that user added today
func (l *PqDB) GetUserTodayTask(ctx context.Context, userID string) (int, error) {
	now := time.Now()
	stmt := `SELECT COUNT(*) FROM tasks WHERE user_id = $1 AND created_date = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, now.Format("2006-01-02"))

	var countTask int
	err := row.Scan(&countTask)

	if err != nil {
		return 0, err
	}

	return countTask, nil
}
