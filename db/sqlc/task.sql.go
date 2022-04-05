// Code generated by sqlc. DO NOT EDIT.
// source: task.sql

package db

import (
	"context"
)

const countTasksCreatedToday = `-- name: CountTasksCreatedToday :one
SELECT
	COUNT(created_at)
FROM
	tasks
WHERE
	OWNER = $1 AND
	created_at :: DATE = NOW() :: DATE
`

func (q *Queries) CountTasksCreatedToday(ctx context.Context, owner string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTasksCreatedToday, owner)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO
	tasks (NAME, OWNER, CONTENT)
VALUES ($1, $2, $3) RETURNING id, name, owner, content, deleted, content_change_at, deleted_at, created_at
`

type CreateTaskParams struct {
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Content string `json:"content"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.Name, arg.Owner, arg.Content)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.Content,
		&i.Deleted,
		&i.ContentChangeAt,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTaskByName = `-- name: DeleteTaskByName :exec
UPDATE
	tasks
SET
	deleted = TRUE,
	deleted_at = NOW()
WHERE
	NAME = $1 AND
	OWNER = $2
`

type DeleteTaskByNameParams struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (q *Queries) DeleteTaskByName(ctx context.Context, arg DeleteTaskByNameParams) error {
	_, err := q.db.ExecContext(ctx, deleteTaskByName, arg.Name, arg.Owner)
	return err
}

const getTask = `-- name: GetTask :one
SELECT
	id, name, owner, content, deleted, content_change_at, deleted_at, created_at
FROM
	tasks
WHERE
	id = $1 AND
	deleted = FALSE
LIMIT
	1
`

func (q *Queries) GetTask(ctx context.Context, id int64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.Content,
		&i.Deleted,
		&i.ContentChangeAt,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getTaskByName = `-- name: GetTaskByName :one
SELECT
	id, name, owner, content, deleted, content_change_at, deleted_at, created_at
FROM
	tasks
WHERE
	NAME = $1 AND
	OWNER = $2 AND
	deleted = FALSE
LIMIT
	1
`

type GetTaskByNameParams struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func (q *Queries) GetTaskByName(ctx context.Context, arg GetTaskByNameParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByName, arg.Name, arg.Owner)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.Content,
		&i.Deleted,
		&i.ContentChangeAt,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listTasksByOwner = `-- name: ListTasksByOwner :many
SELECT
	id, name, owner, content, deleted, content_change_at, deleted_at, created_at
FROM
	tasks
WHERE (
		OWNER = $3 OR
		OWNER = 'admin'
	) AND
	deleted = FALSE
ORDER BY
	id
LIMIT
	$1
OFFSET
	$2
`

type ListTasksByOwnerParams struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Owner  string `json:"owner"`
}

func (q *Queries) ListTasksByOwner(ctx context.Context, arg ListTasksByOwnerParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksByOwner, arg.Limit, arg.Offset, arg.Owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Owner,
			&i.Content,
			&i.Deleted,
			&i.ContentChangeAt,
			&i.DeletedAt,
			&i.CreatedAt,
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

const updateTaskByName = `-- name: UpdateTaskByName :one
UPDATE
	tasks
SET
	CONTENT = $2,
	content_change_at = NOW()
WHERE
	NAME = $1 AND
	OWNER = $3 RETURNING id, name, owner, content, deleted, content_change_at, deleted_at, created_at
`

type UpdateTaskByNameParams struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Owner   string `json:"owner"`
}

func (q *Queries) UpdateTaskByName(ctx context.Context, arg UpdateTaskByNameParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskByName, arg.Name, arg.Content, arg.Owner)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Owner,
		&i.Content,
		&i.Deleted,
		&i.ContentChangeAt,
		&i.DeletedAt,
		&i.CreatedAt,
	)
	return i, err
}
