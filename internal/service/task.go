package service

import (
	"context"
	"time"
	"togo/internal/entity"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, context string, userId int32, createdDate time.Time) (entity.Task, error)
	ListTasks(ctx context.Context, userId int32) ([]entity.Task, error)
	GetTask(ctx context.Context, id int32, userId int32) (entity.Task, error)
	DeleteTask(ctx context.Context, id int32, userId int32) error
	UpdateTask(ctx context.Context, id int32, isDone bool) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) Create(ctx context.Context, content string, userId int32, createdDate time.Time) (entity.Task, error) {
	task, err := t.repo.CreateTask(ctx, content, userId, createdDate)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (t *TaskService) ListTasks(ctx context.Context, userId int32) ([]entity.Task, error) {
	tasks, err := t.repo.ListTasks(ctx, userId)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) GetTask(ctx context.Context, id int32, userId int32) (entity.Task, error) {
	task, err := t.repo.GetTask(ctx, id, userId)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (t *TaskService) DeleteTask(ctx context.Context, id int32, userId int32) error {
	err := t.repo.DeleteTask(ctx, id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskService) UpdateTask(ctx context.Context, id int32, isDone bool) error {
	err := t.repo.UpdateTask(ctx, id, isDone)
	if err != nil {
		return err
	}

	return nil
}
