// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: todo_tasks.sql

package store

import (
	"context"
	"database/sql"
)

const getTaskByID = `-- name: GetTaskByID :one
SELECT
    id, created_at, updated_at, user_id, task_name
FROM
    todo_tasks
WHERE
    id = ?
`

func (q *Queries) GetTaskByID(ctx context.Context, id uint64) (TodoTask, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i TodoTask
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.TaskName,
	)
	return i, err
}

const getTaskByUserID = `-- name: GetTaskByUserID :many
SELECT
    id, created_at, updated_at, user_id, task_name
FROM
    todo_tasks
WHERE
    user_id = ?
`

func (q *Queries) GetTaskByUserID(ctx context.Context, userID uint64) ([]TodoTask, error) {
	rows, err := q.db.QueryContext(ctx, getTaskByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TodoTask
	for rows.Next() {
		var i TodoTask
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.TaskName,
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

const insertTask = `-- name: InsertTask :execresult
INSERT INTO
    todo_tasks(user_id, task_name)
VALUES
    (?, ?)
`

type InsertTaskParams struct {
	UserID   uint64 `json:"user_id"`
	TaskName string `json:"task_name"`
}

func (q *Queries) InsertTask(ctx context.Context, arg InsertTaskParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertTask, arg.UserID, arg.TaskName)
}
