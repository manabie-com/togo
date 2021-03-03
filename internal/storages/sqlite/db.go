package sqlite

import (
	"context"
	"database/sql"

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
func (l *LiteDB) AddTask(ctx context.Context, t *storages.Task) (int64, error) {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) SELECT ?, ?, ?, ? ` +
		`WHERE (select max_todo from users where id = ?) > ` +
		`(select count(id) from tasks where user_id = ? and created_date = ?); ` +
		`SELECT last_insert_rowid();`
	r, err := l.DB.ExecContext(ctx, stmt,
		&t.ID, &t.Content, &t.UserID, &t.CreatedDate,
		&t.UserID,
		&t.UserID, &t.CreatedDate)
	if err != nil {
		return 0, err
	}

	affectedCount, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affectedCount, nil
}

// DeleteTask adds a new task to DB
func (l *LiteDB) DeleteTask(ctx context.Context, t *storages.Task) error {
	stmt := `DELETE FROM main.tasks WHERE main.tasks.id = ?`
	_, err := l.DB.ExecContext(ctx, stmt, &t.ID)
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
