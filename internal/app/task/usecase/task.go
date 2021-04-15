package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/util"
	"github.com/rs/zerolog/log"
	"time"
)

type TaskService struct {
	taskStorage TaskStorage
}

func NewTaskService(taskStorage TaskStorage) *TaskService {
	return &TaskService{taskStorage: taskStorage}
}

//go:generate mockgen -package mock -destination mock/task_mock.go github.com/manabie-com/togo/internal/app/task/usecase TaskStorage
type TaskStorage interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]model.Task, error)
	AddTask(ctx context.Context, task model.Task) error
	LimitReached(ctx context.Context, userID, createdDate string) (bool, error)
}

func (t TaskService) RetrieveTasks(ctx context.Context, createdDate string) ([]model.Task, error) {
	userID, ok := util.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID is not found")
		log.Error().Msg(err.Error())
		return nil, err
	}
	tasks, err := t.taskStorage.RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		log.Error().Str("userID", userID).Str("createdDate", createdDate).Err(err).Msg("unable to retrieve tasks")
		return nil, err
	}
	return tasks, nil
}

func (t TaskService) AddTask(ctx context.Context, task model.Task) (model.Task, error) {
	userID, ok := util.GetUserIDFromContext(ctx)
	if !ok {
		err := errors.New("user ID is not found")
		log.Error().Msg(err.Error())
		return model.Task{}, err
	}
	currentDate := time.Now().Format("2006-01-02")
	task.ID = uuid.New().String()
	task.UserID = userID
	task.CreatedDate = currentDate

	limitIsReached, err := t.taskStorage.LimitReached(ctx, userID, currentDate)
	if err != nil {
		log.Error().Err(err).Msg("unable to check the limit")
		return model.Task{}, err
	}

	if limitIsReached {
		log.Error().Fields(map[string]interface{}{"task": task}).Msg("limit reached")
		return model.Task{}, errors.New("limit reached")
	}

	err = t.taskStorage.AddTask(ctx, task)
	if err != nil {
		log.Error().Fields(map[string]interface{}{"task": task}).Err(err).Msg("unable to add task")
		return model.Task{}, err
	}
	return task, nil
}
