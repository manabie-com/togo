package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"sync"

	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/utils"
)

// TaskUsecase implement usecase of task
type TaskUsecase struct {
	Store *postgres.PostgresDB
}

// DeleteTasks delete tasks if matching userID
func (s *TaskUsecase) DeleteTasks(ctx context.Context, userID sql.NullString) error {
	return s.Store.DeleteTasks(ctx, userID)
}

// ListTasks list tasks if matching userID, createdDate
func (s *TaskUsecase) ListTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	return s.Store.RetrieveTasks(ctx, userID, createdDate)
}

// AddTask list tasks of userID
func (s *TaskUsecase) AddTask(ctx context.Context, task *storages.Task) error {
	return s.Store.AddTask(ctx, task)
}

// CheckMaxTaskPerDay check max task matching userID, createdDate
func (s *TaskUsecase) CheckMaxTaskPerDay(req *http.Request, userID, createdDate string) (bool, error) {
	var err error
	var authUser *storages.User
	var totalTask int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if req.Context().Err() != nil {
			return
		}
		authUser, err = s.Store.GetUser(req.Context(), utils.NullString(userID))
	}()
	go func() {
		defer wg.Done()
		if req.Context().Err() != nil {
			return
		}
		totalTask, err = s.Store.CountTasks(req.Context(), utils.NullString(userID), utils.NullString(createdDate))
	}()
	wg.Wait()

	if err != nil {
		return false, err
	}

	return totalTask >= authUser.MaxTodo, nil
}
