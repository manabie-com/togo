package application

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
	"togo/internal/services/task/domain"
	"togo/internal/services/task/store"
	"togo/internal/services/task/store/postgres"
)

type TaskService struct {
	taskRepo *postgres.TaskRepository
	userRepo *postgres.UserRepository
}

func NewTaskService(taskRepo *postgres.TaskRepository, userRepo *postgres.UserRepository) *TaskService {
	return &TaskService{taskRepo, userRepo}
}

type AddTaskCommand struct {
	UserID      uuid.UUID
	Title       string
	Description string
	DueDate     time.Time
}

func (s *TaskService) CreateTask(cmd AddTaskCommand) (*domain.Task, error) {
	// Find user of the task
	user, err := s.userRepo.FindByID(cmd.UserID)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot find user %s", cmd.UserID)
	}
	if user == nil {
		return nil, errors.Errorf("user %s not found", cmd.UserID)
	}
	// Get user today tasks count
	userTodayTasksCount, err := s.taskRepo.Count(store.CountTasksRequest{
		UserID: &cmd.UserID,
		Day:    &cmd.DueDate,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot count today tasks for user %s", cmd.UserID)
	}
	// Check if user has reached daily task limit
	if userTodayTasksCount >= user.DailyTaskLimit {
		return nil, errors.Errorf("user %s has reached daily task limit", cmd.UserID)
	}
	// Create task
	entity := domain.NewTask(uuid.New(), cmd.UserID, cmd.Title, cmd.Description, cmd.DueDate)
	if err = s.taskRepo.Save(entity); err != nil {
		return nil, errors.Wrap(err, "cannot save task")
	}
	return entity, nil
}
