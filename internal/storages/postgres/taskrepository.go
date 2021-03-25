package postgres

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

const addTask = `
INSERT INTO tasks (id, content, user_id, created_date) VALUES ($1, $2, $3, $4)
RETURNING id
`

// AddTask adds a new task to DB
func (q *Queries) AddTask(ctx context.Context, arg *storages.Task) error {
	row := q.queryRow(ctx, q.addTaskStmt, addTask,
		&arg.ID,
		&arg.Content,
		&arg.UserID,
		&arg.CreatedDate,
	)

	var id string
	err := row.Scan(&id)

	return err
}

const retrieveTasks = `
SELECT id, content, user_id, created_date FROM tasks WHERE user_id = $1 AND created_date = $2
`

type RetrieveTasksParams struct {
	UserID      sql.NullString `json:"user_id"`
	CreatedDate sql.NullString `json:"created_date"`
}

// RetrieveTasks returns tasks if match userID AND createDate.
func (q *Queries) RetrieveTasks(ctx context.Context, arg RetrieveTasksParams) ([]storages.Task, error) {
	rows, err := q.query(ctx, q.retrieveTasksStmt, retrieveTasks, arg.UserID, arg.CreatedDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []storages.Task
	for rows.Next() {
		var i storages.Task
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.UserID,
			&i.CreatedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type CountTaskPerDayParams struct {
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

const countTaskPerDay = `-- name: CountTaskPerDay :one
SELECT count(user_id) FROM tasks WHERE user_id = $1 AND created_date = $2
`

func (q *Queries) CountTaskPerDay(ctx context.Context, arg CountTaskPerDayParams) (int64, error) {
	row := q.queryRow(ctx, q.countTaskPerDayStmt, countTaskPerDay, arg.UserID, arg.CreatedDate)
	var count int64
	err := row.Scan(&count)
	return count, err
}
