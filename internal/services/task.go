package services

import (
	"errors"
	"time"

	"github.com/manabie-com/togo/pkg/model"

	"github.com/manabie-com/togo/internal/storages/postgres"
)

type TaskService interface {
	FindTask(userId int, createdDate *time.Time) ([]model.Task, error)
	CreateTask(user *model.User, content string) error
}

type taskService struct {
	repo postgres.Repository
}

func NewTaskService(repo postgres.Repository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) FindTask(userId int, createdDate *time.Time) ([]model.Task, error) {
	filter := postgres.TaskFilter{
		UserId:      userId,
		CreatedDate: createdDate,
	}

	task, err := s.repo.FindTask(filter)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) CreateTask(user *model.User, content string) error {

	if content == "" {
		return errors.New("empty content")
	}

	t := time.Now()
	tasks, _ := s.FindTask(user.ID, &t)

	if len(tasks) >= user.MaxTodo {
		return errors.New("task limited")
	}

	task := model.Task{
		UserID:      user.ID,
		Content:     content,
		CreatedDate: time.Now(),
	}

	if err := s.repo.SaveTask(&task); err != nil {
		return err
	}

	return nil
}
