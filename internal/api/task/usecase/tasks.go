package usecase

import (
	"context"
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/api/task/storages"
	userStorages "github.com/manabie-com/togo/internal/api/user/storages"
	"github.com/manabie-com/togo/internal/api/utils"
	"github.com/manabie-com/togo/internal/pkg/logger"
)

type Task struct {
	Store           storages.Store
	UserStore       userStorages.Store
	GeneratorUUIDFn utils.GenerateNewUUIDFn
}

func (s *Task) List(ctx context.Context, userID, createdDate string, page, limit int) ([]*storages.Task, error) {
	tasks, err := s.Store.RetrieveTasks(ctx, userID, createdDate, page, limit)
	if err != nil {
		logger.MBErrorf(ctx, "task storage failed to retrieve tasks of user_id %s created_date %s: %v", userID, createdDate, err)
		return nil, errors.New(dictionary.FailedGetRetrieveTasks)
	}

	return tasks, nil
}

func (s *Task) Add(ctx context.Context, userID string, task *storages.Task) (*storages.Task, error) {
	now := utils.GetTimeNowWithDefaultLayoutInString()

	user, err := s.UserStore.Get(ctx, userID)
	if err != nil {
		logger.MBError(ctx, err)
		return nil, errors.New(dictionary.FailedToGetUser)
	}

	tasks, err := s.List(ctx, userID, now, 0, user.MaxTodo+1)
	if err != nil {
		return nil, err
	}

	if len(tasks) >= user.MaxTodo {
		return nil, errors.New(dictionary.UserReachTaskLimit)
	}

	task.UserID = userID
	task.ID = s.GeneratorUUIDFn()
	task.CreatedDate = now

	if err = s.Store.AddTask(ctx, task); err != nil {
		logger.MBErrorln(ctx, err)
		return nil, errors.New(dictionary.StoreTaskFailed)
	}

	return task, nil
}
