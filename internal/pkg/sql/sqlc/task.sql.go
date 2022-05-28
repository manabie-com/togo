// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: task.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const assignTask = `-- name: AssignTask :one
update tasks
SET assignee = $2
WHERE id = $1
returning id, name, assignee, assign_date, description, status, creator, created_at, start_date, end_date
`

type AssignTaskParams struct {
	ID       int32          `json:"id"`
	Assignee sql.NullString `json:"assignee"`
}

func (q *Queries) AssignTask(ctx context.Context, arg *AssignTaskParams) (*Task, error) {
	row := q.queryRow(ctx, q.assignTaskStmt, assignTask, arg.ID, arg.Assignee)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Assignee,
		&i.AssignDate,
		&i.Description,
		&i.Status,
		&i.Creator,
		&i.CreatedAt,
		&i.StartDate,
		&i.EndDate,
	)
	return &i, err
}

const countTaskByAssigneeToday = `-- name: CountTaskByAssigneeToday :one
SELECT count(assign_date)
FROM tasks
where assignee=$1 and DATE(assign_date) = current_date
group BY DATE(assign_date)
`

func (q *Queries) CountTaskByAssigneeToday(ctx context.Context, assignee sql.NullString) (int64, error) {
	row := q.queryRow(ctx, q.countTaskByAssigneeTodayStmt, countTaskByAssigneeToday, assignee)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTask = `-- name: CreateTask :one
insert into tasks
(name, assignee, assign_date, description, status, creator, start_date, end_date)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning id, name, assignee, assign_date, description, status, creator, created_at, start_date, end_date
`

type CreateTaskParams struct {
	Name        string         `json:"name"`
	Assignee    sql.NullString `json:"assignee"`
	AssignDate  time.Time      `json:"assign_date"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	Creator     string         `json:"creator"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
}

func (q *Queries) CreateTask(ctx context.Context, arg *CreateTaskParams) (*Task, error) {
	row := q.queryRow(ctx, q.createTaskStmt, createTask,
		arg.Name,
		arg.Assignee,
		arg.AssignDate,
		arg.Description,
		arg.Status,
		arg.Creator,
		arg.StartDate,
		arg.EndDate,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Assignee,
		&i.AssignDate,
		&i.Description,
		&i.Status,
		&i.Creator,
		&i.CreatedAt,
		&i.StartDate,
		&i.EndDate,
	)
	return &i, err
}

const deleteTask = `-- name: DeleteTask :exec
delete
from tasks
where id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :one
select id, name, assignee, assign_date, description, status, creator, created_at, start_date, end_date
from tasks
where id = $1
limit 1
`

func (q *Queries) GetTask(ctx context.Context, id int32) (*Task, error) {
	row := q.queryRow(ctx, q.getTaskStmt, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Assignee,
		&i.AssignDate,
		&i.Description,
		&i.Status,
		&i.Creator,
		&i.CreatedAt,
		&i.StartDate,
		&i.EndDate,
	)
	return &i, err
}

const listTasks = `-- name: ListTasks :many
select id, name, assignee, assign_date, description, status, creator, created_at, start_date, end_date
from tasks
order by created_at
limit $1 offset $2
`

type ListTasksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTasks(ctx context.Context, arg *ListTasksParams) ([]*Task, error) {
	rows, err := q.query(ctx, q.listTasksStmt, listTasks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Assignee,
			&i.AssignDate,
			&i.Description,
			&i.Status,
			&i.Creator,
			&i.CreatedAt,
			&i.StartDate,
			&i.EndDate,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
