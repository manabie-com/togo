package usecase

import (
	"context"

	"github.com/manabie-com/togo/internal/model"
)

type TaskService interface {
	ListTasks(ctx context.Context, userId string, createdDate string) (res []*model.Task, err error)
	AddTask(ctx context.Context, t *model.Task) error
	IsAllowedToAddTask(ctx context.Context, userId string) bool
}

type taskService struct {
	taskRespository model.TaskRespository
}

func NewTaskService(tr model.TaskRespository) TaskService {
	return &taskService{
		taskRespository: tr,
	}
}

func (ts *taskService) ListTasks(ctx context.Context, userId string, createdDate string) (res []*model.Task, err error) {
	tasks, err := ts.taskRespository.RetrieveTasks(
		ctx, userId, createdDate,
	)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (ts *taskService) AddTask(ctx context.Context, t *model.Task) error {

	err := ts.taskRespository.AddTask(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (ts *taskService) IsAllowedToAddTask(ctx context.Context, userId string) bool {
	return ts.taskRespository.IsAllowedToAddTask(ctx, userId)
}
