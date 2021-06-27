package services

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/manabie-com/togo/internal/tokens"
	"strconv"

	"github.com/manabie-com/togo/internal/repositories"

)

const (
	cachedAddTaskTimesKeyPrefix = "add_task_times_%s"
	taskPerDayExpiredTime = 86400
	noChangeExpiredTime = -1
)

type TodoService interface {
	ListTasks(userID, createdAt string) (*[]repositories.Task, error)
	AddTask(userID string, task *repositories.Task) (*repositories.Task, error)
	GetAuthToken(userID, password string) (string, error)
	ValidToken(token string) (userID string, valid bool)
}

type ToDoServiceImpl struct {
	TaskRepo     repositories.TaskRepo
	TokenManager tokens.TokenManager
	CachingRepo  repositories.CachingRepo
	UserRepo     repositories.UserRepo
}

func NewToDoService(taskRepo repositories.TaskRepo, userRepo repositories.UserRepo, tokenManager tokens.TokenManager, cachingRepo repositories.CachingRepo) *ToDoServiceImpl {
	return &ToDoServiceImpl{
		TaskRepo:     taskRepo,
		UserRepo:     userRepo,
		TokenManager: tokenManager,
		CachingRepo:  cachingRepo,
	}
}


func (s *ToDoServiceImpl) ListTasks(userID, createdAt string) (*[]repositories.Task, error) {
	return s.TaskRepo.ListTask(userID, createdAt)
}

func (s *ToDoServiceImpl) AddTask(userID string, task *repositories.Task) (*repositories.Task, error) {
	maxToDo, err := s.UserRepo.GetMaxToDoOfUser(userID)
	if err != nil {
		return nil, err
	}

	cachedAddTaskTimesKey := buildCachedAddTaskTimesKey(userID)
	addTaskTimes, err := s.CachingRepo.Get(buildCachedAddTaskTimesKey(userID))

	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
		err := s.CachingRepo.Increase(cachedAddTaskTimesKey, taskPerDayExpiredTime)
		if err != nil {
			return nil, err
		}
	} else {
		counter, err := strconv.Atoi(addTaskTimes)
		if err != nil {
			return nil, err
		}

		if counter >= maxToDo {
			return nil, errors.New("exceed the limited times to add task")
		}
		err = s.CachingRepo.Increase(cachedAddTaskTimesKey, noChangeExpiredTime)
		if err != nil {
			return nil, err
		}
	}
	return s.TaskRepo.AddTask(task)
}

func buildCachedAddTaskTimesKey(userID string) string  {
	return fmt.Sprintf(cachedAddTaskTimesKeyPrefix, userID)
}

func (s *ToDoServiceImpl) GetAuthToken(userID, password string) (string, error) {
	return s.TokenManager.GetAuthToken(userID, password)
}

func (s *ToDoServiceImpl) ValidToken(token string) (userID string, valid bool) {
	return s.TokenManager.ValidToken(token)
}

