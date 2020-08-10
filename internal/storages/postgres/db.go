package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/manabie-com/togo/internal/storages"
)

// PgDB for working with postgres
type PgDB struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (l *PgDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
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

// CountTaskToday returns number of user task today.
func (l *PgDB) CountTaskToday(ctx context.Context, userID sql.NullString) (int, error) {
	stmt := `SELECT count(id) FROM tasks WHERE user_id = $1 AND created_date = to_char(now(),'YYYY-MM-DD')`
	row := l.DB.QueryRowContext(ctx, stmt, userID)

	countTask := 0
	if err := row.Scan(&countTask); err != nil {
		return 0, err
	}

	return countTask, nil
}

// AddTask adds a new task to DB
func (l *PgDB) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

// ValidateUser returns tasks if match userID AND password
func (l *PgDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = $1 AND password = $2`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		fmt.Println("err:", err)
		return false
	}

	return true
}
