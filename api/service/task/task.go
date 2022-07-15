package task

import (
	"context"

	"manabie/todo/models"
	"manabie/todo/repository/task"
)

type TaskService interface {
	Index(ctx context.Context) ([]*models.Task, error)
	Show(ctx context.Context) (*models.Task, error)
	Create(ctx context.Context, t *models.Task) error
	Update(ctx context.Context, t *models.Task) error
	Delete(ctx context.Context, t *models.Task) error
}

type service struct{}

func NewTaskService(tr task.TaskRespository) TaskService {
	return &service{}
}

func (s *service) Index(ctx context.Context) ([]*models.Task, error) {
	return nil, nil
}

func (s *service) Show(ctx context.Context) (*models.Task, error) {
	return nil, nil
}

func (s *service) Create(ctx context.Context, t *models.Task) error {
	return nil
}

func (s *service) Update(ctx context.Context, t *models.Task) error {
	return nil
}

func (s *service) Delete(ctx context.Context, t *models.Task) error {
	return nil
}
