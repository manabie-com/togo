package handler

import (
	"context"
	"time"

	"togo/common"
	"togo/internal/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, context string, userID int32, createdDate time.Time) (*entity.Task, error)
	ListTasks(ctx context.Context, userId int32, isDone bool, createdDate time.Time) ([]entity.Task, error)
	GetTask(ctx context.Context, id int32, userId int32) (*entity.Task, error)
	DeleteTask(ctx context.Context, id int32, userId int32) error
	UpdateTask(ctx context.Context, id int32, isDone bool) (*entity.Task, error)
	CountTaskByUser(ctx context.Context, userID int32) (int32, error)
}

type TaskRedisRepo interface {
	GetTask(ctx context.Context, id int32) (*entity.Task, error)
	SetTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int32) error
}

type TaskHandler struct {
	repo    TaskRepository
	rdbRepo TaskRedisRepo
}

func NewTaskHandler(repo TaskRepository, rdbRepo TaskRedisRepo) *TaskHandler {
	return &TaskHandler{repo: repo, rdbRepo: rdbRepo}
}

func (t *TaskHandler) CountTaskByUser(ctx context.Context, userID int32) (int32, error) {
	count, err := t.repo.CountTaskByUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t *TaskHandler) CreateTask(ctx context.Context, content string, userID int32, createdDate time.Time) (*entity.Task, error) {
	task, err := t.repo.CreateTask(ctx, content, userID, createdDate)
	if err != nil {
		return nil, err
	}

	_ = t.rdbRepo.SetTask(ctx, task)

	return task, nil
}

func (t *TaskHandler) ListTask(ctx context.Context, userID int32, isDone bool, createdDate time.Time) ([]entity.Task, error) {
	tasks, err := t.repo.ListTasks(ctx, userID, isDone, createdDate)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskHandler) GetTask(ctx context.Context, id int32, userID int32) (*entity.Task, error) {
	task, err := t.rdbRepo.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	if task != nil {
		if task.UserID == userID {
			return task, nil
		}

		return nil, common.ErrTaskNotFound
	}

	task, err = t.repo.GetTask(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	_ = t.rdbRepo.SetTask(ctx, task)

	return task, nil
}

func (t *TaskHandler) DeleteTask(ctx context.Context, id int32, userID int32) error {
	if err := t.repo.DeleteTask(ctx, id, userID); err != nil {
		return err
	}

	_ = t.rdbRepo.DeleteTask(ctx, id)

	return nil
}

func (t *TaskHandler) UpdateTask(ctx context.Context, id int32, isDone bool) error {
	task, err := t.repo.UpdateTask(ctx, id, isDone)
	if err != nil {
		return err
	}

	_ = t.rdbRepo.SetTask(ctx, task)

	return nil
}
