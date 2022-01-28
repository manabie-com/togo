package service

import (
	"context"
	"errors"

	"github.com/manabie-com/togo/pkg/errorx"

	"gorm.io/gorm"

	"github.com/manabie-com/togo/model"
	intCtx "github.com/manabie-com/togo/pkg/ctx"

	"github.com/manabie-com/togo/internal/task/repository"
)

type TaskService interface {
	CreateTask(context.Context, *CreateTaskArgs) error
	UpdateTask(context.Context, *UpdateTaskArgs) error
	DeleteTask(context.Context, *DeleteTaskArgs) error
	GetTasks(context.Context, *GetTasksArgs) ([]*Task, error)
	GetTask(context.Context, *GetTaskArgs) (*Task, error)
}

type taskService struct {
	db       *gorm.DB
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository, db *gorm.DB) TaskService {
	return &taskService{
		db:       db,
		taskRepo: taskRepo,
	}
}

func (t *taskService) GetTasks(ctx context.Context, args *GetTasksArgs) ([]*Task, error) {
	currentUser := intCtx.Get(ctx, intCtx.UserKey).(*model.User)

	tasks, err := t.taskRepo.GetTasks(ctx, &repository.GetTasksQuery{
		UserID: currentUser.ID,
		Limit:  args.Limit,
		Offset: args.Offset,
	})
	if err != nil {
		return nil, err
	}
	return convertModelTasksToServiceTasks(tasks), nil
}

func (t *taskService) GetTask(ctx context.Context, args *GetTaskArgs) (*Task, error) {
	currentUser := intCtx.Get(ctx, intCtx.UserKey).(*model.User)
	task, err := t.taskRepo.GetTask(ctx, &model.Task{
		ID:     args.ID,
		UserID: currentUser.ID,
	})
	if err != nil {
		return nil, err
	}
	return convertModelTaskToServiceTask(task), nil
}

func (t *taskService) DeleteTask(ctx context.Context, args *DeleteTaskArgs) error {
	currentUser := intCtx.Get(ctx, intCtx.UserKey).(*model.User)
	task, err := t.taskRepo.GetTask(ctx, &model.Task{
		ID:     args.ID,
		UserID: currentUser.ID,
	})
	if err != nil {
		return err
	}
	tx := t.db.Begin()
	if err := t.taskRepo.DeleteTask(tx, task); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *taskService) CreateTask(ctx context.Context, args *CreateTaskArgs) error {
	currentUser := intCtx.Get(ctx, intCtx.UserKey).(*model.User)
	taskCount, err := t.taskRepo.CountByUserID(ctx, currentUser.ID)
	if err != nil {
		return err
	}
	if int(taskCount) >= currentUser.LimitTask {
		return errorx.ErrInternal(errors.New("Can not reach out task limit"))
	}
	tx := t.db.Begin()
	if err := t.taskRepo.SaveTask(tx, &model.Task{
		Content: args.Content,
		UserID:  currentUser.ID,
	}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *taskService) UpdateTask(ctx context.Context, args *UpdateTaskArgs) error {
	currentUser := intCtx.Get(ctx, intCtx.UserKey).(*model.User)
	task, err := t.taskRepo.GetTask(ctx, &model.Task{
		ID:     args.TaskID,
		UserID: currentUser.ID,
	})
	if err != nil {
		return err
	}

	tx := t.db.Begin()
	if err := t.taskRepo.UpdateTask(tx, &model.Task{
		ID:      task.ID,
		Content: args.Content,
	}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
