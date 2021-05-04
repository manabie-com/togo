package services

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/models"
	pkg "github.com/manabie-com/togo/internal/pkg/utils"
	"github.com/manabie-com/togo/internal/repositories"
	"strings"
	"time"
)

// ToDoService implement HTTP server

type ToDoService interface {
	GetAuthToken(ctx context.Context, id, password string) bool
	ListTasks(ctx context.Context, userId, createdDate string) ([]*models.Task, error)
	AddTask(ctx context.Context, t *models.Task) error
}

func NewToDoService(repo repositories.TaskRepo) ToDoService {
	return &toDoService{Repo: repo}
}

type toDoService struct {
	utils  pkg.Utils
	JWTKey string
	Repo   repositories.TaskRepo
}

func (s *toDoService) GetAuthToken(ctx context.Context, id, password string) bool {
	if len(strings.TrimSpace(id)) == 0 || len(strings.TrimSpace(password)) == 0 {
		return false
	}
	existUser := s.Repo.ValidateUser(ctx, id, password)
	return existUser
}

func (s *toDoService) ListTasks(ctx context.Context, userId, createdDate string) ([]*models.Task, error) {
	tasks, err := s.Repo.RetrieveTasks(
		ctx,
		userId,
		createdDate,
	)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *toDoService) AddTask(ctx context.Context, t *models.Task) error {
	now := time.Now()
	maxToDo, err := s.Repo.GetMaxTaskPerDay(ctx, t.UserID)
	if err != nil {
		return err
	}
	checkAddTask := s.Repo.CheckTaskPerDayOfAnUser(ctx, maxToDo, t.UserID, now.Format("2006-01-02"))
	if !checkAddTask {
		return errors.New("max task per day")
	}
	err = s.Repo.AddTask(ctx, t)
	if err != nil {
		return err
	}
	return nil
}
