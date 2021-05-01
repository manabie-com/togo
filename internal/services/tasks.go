package services

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
)

// ErrUserReachDailyRequestLimit returns if a user has reached daily request limit
var ErrUserReachDailyRequestLimit = errors.New("User has reached daily request limit")

// ToDoService contains methods to handle users' todos related requests
type ToDoService interface {
	// ValidateUser checks if a user existed
	ValidateUser(ctx context.Context, userID, password string) bool
	// ListTasks returns tasks from storage
	ListTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error)
	// AddTask adds a task to storage
	AddTask(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error)
}

// todoService implement ToDoService
type todoService struct {
	Store                  storages.Repository
	mux                    sync.RWMutex
	lastUsersCreatedDate   map[string]string
	usersDailyRequestCount map[string]int
}

// NewToDoService returns a ToDoService
func NewToDoService(store storages.Repository) ToDoService {
	return &todoService{
		Store:                  store,
		lastUsersCreatedDate:   make(map[string]string),
		usersDailyRequestCount: make(map[string]int),
	}
}

func (s *todoService) ValidateUser(ctx context.Context, userID, password string) bool {
	ok := s.Store.ValidateUser(ctx, userID, password)
	return ok
}

func (s *todoService) ListTasks(ctx context.Context, userID, createdDate string) ([]*storages.Task, error) {
	tasks, err := s.Store.RetrieveTasks(ctx, userID, createdDate)
	return tasks, err
}

func (s *todoService) AddTask(ctx context.Context, userID string, reqBody []byte) (*storages.Task, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	t := &storages.Task{}
	err := json.Unmarshal(reqBody, &t)
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")
	if _, ok := s.lastUsersCreatedDate[userID]; !ok {
		s.lastUsersCreatedDate[userID] = today
	}
	if _, ok := s.usersDailyRequestCount[userID]; !ok || today != s.lastUsersCreatedDate[userID] {
		if ok {
			s.usersDailyRequestCount[userID], err = s.Store.MaxTodo(ctx, userID)
			if err != nil {
				return nil, err
			}
		} else {
			// In case of failure and the service needs to be restarted, subtract the user's today
			// requests already made (in db) from the user's today total request limit
			todayRequestCountFromDB, err := s.Store.LoadTasksCount(ctx, userID, today)
			if err != nil {
				return nil, err
			}
			userDailyMaxTodo, err := s.Store.MaxTodo(ctx, userID)
			if err != nil {
				return nil, err
			}
			s.usersDailyRequestCount[userID] = userDailyMaxTodo - todayRequestCountFromDB
		}

	}
	if s.usersDailyRequestCount[userID] <= 0 {
		return nil, ErrUserReachDailyRequestLimit
	}
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = today

	err = s.Store.AddTask(ctx, t)
	if err != nil {
		return nil, err
	}

	s.usersDailyRequestCount[userID]--
	s.lastUsersCreatedDate[userID] = today

	return t, nil
}
