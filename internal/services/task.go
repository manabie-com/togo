package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	httpPkg "github.com/manabie-com/togo/pkg/http"
	"github.com/manabie-com/togo/pkg/transactional"
)

type TaskService interface {
	ListTasks(ctx context.Context, createdDate time.Time) ([]models.Task, error)
	AddTask(ctx context.Context, task *models.Task) (*models.Task, error)
}

type taskService struct {
	db   transactional.DB
	repo *repositories.Repository
}

func newTaskService(repo *repositories.Repository, db transactional.DB) TaskService {
	return &taskService{
		db:   db,
		repo: repo,
	}
}

func (s *taskService) ListTasks(ctx context.Context, createdDate time.Time) ([]models.Task, error) {
	userID, ok := ctx.Value(httpPkg.UserIDKey).(string)
	if !ok || userID == "" {
		return nil, errors.New("not authorize")
	}
	tasks, err := s.repo.TaskRepository.ListTasks(ctx, userID, createdDate)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) AddTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	userID, ok := ctx.Value(httpPkg.UserIDKey).(string)
	if !ok || userID == "" {
		return nil, errors.New("not authorize")
	}
	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = time.Now()
	var resp *models.Task

	if err := transactional.WithTx(ctx, s.db, func(ctx context.Context) (err error) {
		err = s.repo.TaskRepository.AddTask(ctx, *task)
		if err != nil {
			return err
		}
		resp, err = s.repo.TaskRepository.GetTask(ctx, task.ID)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return resp, nil
}
