// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const countUserTasks = `-- name: CountUserTasks :one
SELECT count(*) FROM tasks WHERE user_id = $1 AND created_at BETWEEN NOW() - INTERVAL '24 HOURS' AND NOW()
`

func (q *Queries) CountUserTasks(ctx context.Context, userID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUserTasks, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (title, content, is_complete, user_id) VALUES ($1, $2, $3, $4) RETURNING id, title, content, is_complete, user_id, created_at
`

type CreateTaskParams struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	IsComplete bool   `json:"is_complete"`
	UserID     int64  `json:"user_id"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.Title,
		arg.Content,
		arg.IsComplete,
		arg.UserID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.IsComplete,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, daily_task_limit) VALUES ($1, $2) RETURNING id, username, daily_task_limit, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	DailyTaskLimit int32  `json:"daily_task_limit"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.DailyTaskLimit)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.DailyTaskLimit,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, username, daily_task_limit, created_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.DailyTaskLimit,
		&i.CreatedAt,
	)
	return i, err
}
