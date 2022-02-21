package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/chi07/todo/internal/model"
)

type CreateTaskService struct {
	taskRepo       CreateTaskRepo
	limitationRepo LimitationRepo
}

func NewCreateTaskService(taskRepo CreateTaskRepo, limitationRepo LimitationRepo) *CreateTaskService {
	return &CreateTaskService{
		taskRepo:       taskRepo,
		limitationRepo: limitationRepo,
	}
}

func (s *CreateTaskService) CreateTask(ctx context.Context, task *model.Task) (uuid.UUID, error) {
	// Check maximum limit of user
	userID := task.UserID
	limitation, err := s.limitationRepo.GetByUserID(ctx, userID)
	if err != nil {
		return uuid.Nil, err
	}

	if limitation == nil {
		return uuid.Nil, model.ErrorNotFound
	}

	// check user can create a task
	noOfTasks, err := s.taskRepo.CountUserTasks(ctx, userID)
	if err != nil {
		return uuid.Nil, err
	}

	if noOfTasks >= limitation.LimitTask {
		return uuid.Nil, model.ErrorNotAllowed
	}

	// set logic for new task
	now := time.Now()
	task.ID = uuid.New()
	task.Status = model.StatusNew
	task.Priority = model.PrioryNormal
	task.CreatedAt = now
	task.UpdatedAt = now

	taskID, err := s.taskRepo.Create(ctx, task)

	if err != nil || taskID == uuid.Nil {
		return uuid.Nil, errors.Wrap(err, "cannot create task")
	}
	return taskID, nil
}
