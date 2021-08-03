package domain

import (
	"context"
	"time"

	"togo/common"
	"togo/internal/entity"
)

type TaskHandler interface {
	CountTaskByUser(ctx context.Context, userID int32) (int32, error)
	CreateTask(ctx context.Context, content string, userID int32, createdDate time.Time) (*entity.Task, error)
	ListTask(ctx context.Context, userID int32, isDone bool, createdDate time.Time) ([]entity.Task, error)
	GetTask(ctx context.Context, id int32, userID int32) (*entity.Task, error)
	DeleteTask(ctx context.Context, id int32, userID int32) error
	UpdateTask(ctx context.Context, id int32, isDone bool) error
}

type TaskDomain struct {
	handler TaskHandler
}

func NewTaskDomain(handler TaskHandler) *TaskDomain {
	return &TaskDomain{handler: handler}
}

func (t *TaskDomain) Create(ctx context.Context, user *entity.User, content string, createdDate time.Time) (*entity.Task, error) {
	count, err := t.handler.CountTaskByUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if count >= user.MaxTodo {
		return nil, common.ErrTooManyTask
	}

	task, err := t.handler.CreateTask(ctx, content, user.ID, createdDate)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskDomain) ListTasks(ctx context.Context, userID int32, isDone bool, createdDate time.Time) ([]entity.Task, error) {
	tasks, err := t.handler.ListTask(ctx, userID, isDone, createdDate)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskDomain) GetTask(ctx context.Context, id int32, userID int32) (*entity.Task, error) {
	task, err := t.handler.GetTask(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskDomain) DeleteTask(ctx context.Context, id int32, userID int32) error {
	_, err := t.handler.GetTask(ctx, id, userID)
	if err != nil {
		return err
	}

	if err = t.handler.DeleteTask(ctx, id, userID); err != nil {
		return err
	}

	return nil
}

func (t *TaskDomain) UpdateTask(ctx context.Context, userID int32, id int32, isDone bool) error {
	_, err := t.handler.GetTask(ctx, id, userID)
	if err != nil {
		return err
	}

	if err = t.handler.UpdateTask(ctx, id, isDone); err != nil {
		return err
	}

	return nil
}
