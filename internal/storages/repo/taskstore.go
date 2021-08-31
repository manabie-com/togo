package repo

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

// TaskStore for working with tasks table
type TaskStore struct {
	DB *sql.DB
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (this *TaskStore) RetrieveTasks(ctx context.Context, userID string, createdDate sql.NullString) ([]*storages.Task, error) {
	stmt := `SELECT id
				  , content
				  , user_id
				  , created_date
			   FROM tasks
			  WHERE user_id = $1
			    AND created_date = $2`
	rows, err := this.DB.QueryContext(ctx, stmt, userID, createdDate)
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

// CountTask return number of tasks by given date
func (this *TaskStore) CountTask(ctx context.Context, userID, createdDate string) (int, error) {
	stmt := `SELECT count(*) AS cnt
			   FROM tasks
			  WHERE user_id = $1
			    AND created_date = $2`
	row := this.DB.QueryRowContext(ctx, stmt, userID, createdDate)

	var count int = 0
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// AddTask adds a new task to DB
func (this *TaskStore) AddTask(ctx context.Context, t *storages.Task) error {
	stmt := `INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
	_, err := this.DB.ExecContext(ctx, stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}
