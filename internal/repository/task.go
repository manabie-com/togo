package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"togo/internal/entity"
	"togo/internal/postgresql"
)

func (r *Repo) CreateTask(ctx context.Context, context string, userId int32, createdDate time.Time) (entity.Task, error) {
	task, err := r.q.InsertTask(ctx, postgresql.InsertTaskParams{
		Content:     context,
		UserID:      userId,
		CreatedDate: createdDate,
	})

	if err != nil {
		return entity.Task{}, err
	}

	return task.MapToEntity(), nil
}

func (r *Repo) ListTasks(ctx context.Context, userId int32, isDone bool, createdDate time.Time) ([]entity.Task, error) {
	tasks, err := r.q.ListTasks(ctx, postgresql.ListTasksParams{
		UserID:      userId,
		IsDone:      isDone,
		CreatedDate: createdDate,
	})
	if err != nil {
		return nil, err
	}

	res := make([]entity.Task, len(tasks))

	for i, task := range tasks {
		res[i] = task.MapToEntity()
	}

	return res, nil
}

func (r *Repo) GetTask(ctx context.Context, id int32, userId int32) (entity.Task, error) {
	task, err := r.q.GetTask(ctx, postgresql.GetTaskParams{
		ID:     id,
		UserID: userId,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Task{}, errors.New("No record")
		}

		return entity.Task{}, err
	}

	return task.MapToEntity(), nil
}

func (r *Repo) DeleteTask(ctx context.Context, id int32, userId int32) error {
	err := r.q.DeleteTask(ctx, postgresql.DeleteTaskParams{
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

func (r *Repo) UpdateTask(ctx context.Context, id int32, isDone bool) error {
	err := r.q.UpdateTask(ctx, postgresql.UpdateTaskParams{
		ID: id, IsDone: isDone,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CountTaskByUser(ctx context.Context, userID int32) (int32, error) {
	count, err := r.q.CountTaskByUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	return int32(count), nil
}
