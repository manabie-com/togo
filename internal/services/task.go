package services

import (
	"context"
	"errors"
	"time"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	httpPkg "github.com/manabie-com/togo/pkg/http"
	"github.com/manabie-com/togo/pkg/txmanager"
)

type TaskService interface {
	ListTasks(ctx context.Context, createdDate time.Time) ([]models.Task, error)
	AddTask(ctx context.Context, task *models.Task) (*models.Task, error)
}

type taskService struct {
	repo *repositories.Repository
	tx   txmanager.TransactionManager
}

func newTaskService(repo *repositories.Repository, tx txmanager.TransactionManager) TaskService {
	return &taskService{
		repo,
		tx,
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
	task.UserID = userID
	var (
		resp   *models.Task
		taskID string
		err    error
	)
	//
	tx := s.tx.Begin(ctx)
	ctx = tx.InjectTransaction(ctx)
	defer func() {
		tx.End(ctx, err)
	}()
	defer tx.Recover(ctx)

	taskID, err = s.repo.TaskRepository.AddTask(ctx, *task)
	if err != nil {
		return nil, err
	}
	if taskID == "" {
		return nil, errors.New("limit 5 tasks per day")
	}

	resp, err = s.repo.TaskRepository.GetTask(ctx, taskID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
