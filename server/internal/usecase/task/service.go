package task

import (
	"context"
	taskEntity "github.com/HoangVyDuong/togo/internal/storages/task"
	"github.com/HoangVyDuong/togo/pkg/define"
	"github.com/HoangVyDuong/togo/pkg/logger"
	"time"
)

type taskService struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) Service {
	return &taskService{
		repo: r,
	}
}


func (ts *taskService) GetTasks(ctx context.Context, userId uint64) ([]taskEntity.Task, error) {
	if userId == 0 {
		logger.Errorf("[TaskService][GetTasks] param invalid. UserID: %d", userId)
		return nil, define.FailedValidation
	}
	return ts.repo.RetrieveTasks(ctx, userId)
}

func (ts *taskService) CreateTask(ctx context.Context, taskEntity taskEntity.Task) error {
	if taskEntity.UserID <= 0 || taskEntity.Content == "" || taskEntity.ID <= 0 {
		logger.Errorf("[TaskService][CreateTask] param invalid. UserID: %d, taskID: %d", taskEntity.UserID, taskEntity.ID)
		return define.FailedValidation
	}

	err := ts.repo.AddTask(ctx, taskEntity, time.Now().UTC())
	return err
}

func (ts *taskService) DeleteTask(ctx context.Context, taskId uint64) error {
	if taskId <= 0 {
		logger.Errorf("[TaskService][DeleteTask] param invalid. TaskID: %d", taskId)
		return define.FailedValidation
	}

	return ts.repo.SoftDeleteTask(ctx, taskId, time.Now().UTC())
}