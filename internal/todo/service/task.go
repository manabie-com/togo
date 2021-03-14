package service

import (
	"time"

	"github.com/manabie-com/togo/internal/pkg/config"
	d "github.com/manabie-com/togo/internal/todo/domain"
)

type TaskService struct {
	TaskRepo d.TaskRepository
}

func NewTaskService(taskRepo d.TaskRepository) *TaskService {
	return &TaskService{taskRepo}
}

func (s *TaskService) ListTaskForUser(userID int, dateStr string) ([]*d.Task, error) {
	tasks, err := s.TaskRepo.GetTasksForUser(userID, dateStr)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) CreateTaskForUser(userID int, param d.TaskCreateParam) (*d.Task, error) {
	numTaskToday, err := s.TaskRepo.GetTaskCount(userID, time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, err
	}

	if numTaskToday >= config.GetEnvInt("MAX_TASKS_DAILY") {
		return nil, d.ErrTaskLimitReached
	}

	task, err := s.TaskRepo.CreateTaskForUser(userID, param)
	if err != nil {
		return nil, err
	}

	return task, nil
}
