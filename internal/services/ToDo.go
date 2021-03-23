package services

import (
	"context"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/core"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
)

// ToDoService implement HTTP server
type ToDoService struct {
	UserRepository storages.UserRepository
	TaskRepository storages.TaskRepository
}

func (s *ToDoService) ValidateUser(ctx context.Context, userID string, password string) bool {
	return s.UserRepository.ValidateUser(ctx, userID, password)
}

func (s *ToDoService) ListTasks(ctx context.Context, userID string, createdDate string) ([]*entities.Task, error) {
	return s.TaskRepository.RetrieveTasks(
		ctx,
		userID,
		createdDate,
	)
}

func (s *ToDoService) AddTask(ctx context.Context, userID string, t *entities.Task) error {
	err := s.isAllowToAddTask(ctx, userID)

	if err != nil {
		return err
	}

	now := time.Now()
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = utils.FormatTimeToString(now)

	return s.TaskRepository.AddTask(ctx, t)
}

//TODO Lock for concurrency request by redis.
func (s *ToDoService) isAllowToAddTask(ctx context.Context, userID string) error {

	taskPerDay, err := s.TaskRepository.CountTaskPerDayByUserID(ctx, userID)

	if err != nil {
		return err
	}

	if taskPerDay >= uint(config.LimitAllowTaskPerDay) {
		return core.ERROR_EXCEED_TASK_LIMITS
	}

	return nil
}
