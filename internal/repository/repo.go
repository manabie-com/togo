package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"togo/internal/entity"
	"togo/internal/postgresql"
)

type Repo struct {
	q *postgresql.Queries
	*sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		q: postgresql.New(db),
	}
}

func (t *Repo) CreateTask(ctx context.Context, context string, userId int32, createdDate time.Time) (entity.Task, error) {
	task, err := t.q.InsertTask(ctx, postgresql.InsertTaskParams{
		Content:     context,
		UserID:      userId,
		CreatedDate: createdDate,
	})

	if err != nil {
		return entity.Task{}, err
	}

	return task.MapToEntity(), nil
}

func (t *Repo) ListTasks(ctx context.Context, userId int32) ([]entity.Task, error) {
	tasks, err := t.q.ListTasks(ctx, userId)
	if err != nil {
		return nil, err
	}

	res := make([]entity.Task, len(tasks))

	for i, task := range tasks {
		res[i] = task.MapToEntity()
	}

	return res, nil
}

func (t *Repo) GetTask(ctx context.Context, id int32, userId int32) (entity.Task, error) {
	task, err := t.q.GetTask(ctx, postgresql.GetTaskParams{
		ID:     id,
		UserID: userId,
	})

	if err != nil {
		return entity.Task{}, err
	}

	return task.MapToEntity(), nil
}

func (t *Repo) DeleteTask(ctx context.Context, id int32, userId int32) error {
	err := t.q.DeleteTask(ctx, postgresql.DeleteTaskParams{
		ID: id, UserID: userId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("task not found")
		}
		return err
	}

	return nil
}

func (t *Repo) UpdateTask(ctx context.Context, id int32, isDone bool) error {
	err := t.q.UpdateTask(ctx, postgresql.UpdateTaskParams{
		ID: id, IsDone: isDone,
	})
	if err != nil {
		return err
	}

	return nil
}
