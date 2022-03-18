package service

import (
	"context"
	"errors"

	repository2 "github.com/manabie-com/togo/internal/user/repository"

	"github.com/manabie-com/togo/pkg/errorx"

	"gorm.io/gorm"

	"github.com/manabie-com/togo/internal/task/repository"
	"github.com/manabie-com/togo/model"
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
	userRepo repository2.UserRepository
}

func NewTaskService(userRepo repository2.UserRepository, taskRepo repository.TaskRepository, db *gorm.DB) TaskService {
	return &taskService{
		db:       db,
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

func (t *taskService) GetTasks(ctx context.Context, args *GetTasksArgs) ([]*Task, error) {
	user, err := t.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return nil, err
	}

	tasks, err := t.taskRepo.GetTasks(ctx, &repository.GetTasksQuery{
		UserID: user.ID,
		Limit:  args.Limit,
		Offset: args.Offset,
	})
	if err != nil {
		return nil, err
	}
	return convertModelTasksToServiceTasks(tasks), nil
}

func (t *taskService) GetTask(ctx context.Context, args *GetTaskArgs) (*Task, error) {
	user, err := t.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return nil, err
	}
	task, err := t.taskRepo.GetTask(ctx, &model.Task{
		ID:     args.ID,
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}
	return convertModelTaskToServiceTask(task), nil
}

func (t *taskService) DeleteTask(ctx context.Context, args *DeleteTaskArgs) error {
	user, err := t.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return err
	}
	task, err := t.taskRepo.GetTask(ctx, &model.Task{
		ID:     args.ID,
		UserID: user.ID,
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
	if args.Content == "" || args.UserID == 0 {
		return errorx.ErrInvalidParameter(errors.New("Missing arguments"))
	}
	user, err := t.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return err
	}
	taskCount, err := t.taskRepo.CountByUserID(ctx, user.ID)
	if err != nil {
		return err
	}
	if int(taskCount) >= user.TaskLimit {
		return errorx.ErrInternal(errors.New("Can not reach out task limit today"))
	}
	tx := t.db.Begin()
	if err := t.taskRepo.SaveTask(tx, &model.Task{
		Content: args.Content,
		UserID:  args.UserID,
	}); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *taskService) UpdateTask(ctx context.Context, args *UpdateTaskArgs) error {
	user, err := t.userRepo.GetUser(ctx, &model.User{
		ID: args.UserID,
	})
	if err != nil {
		return err
	}
	task, err := t.taskRepo.GetTask(ctx, &model.Task{
		ID:     args.TaskID,
		UserID: user.ID,
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
