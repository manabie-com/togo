package domain

import (
	"context"
	"time"

	"github.com/perfectbuii/togo/common/errors"
	"github.com/perfectbuii/togo/internal/storages"
	"github.com/perfectbuii/togo/utils"
)

type TaskDomain interface {
	GetList(ctx context.Context, createdDate string) ([]*storages.Task, error)
	Create(ctx context.Context, content string) (*storages.Task, error)
}

type taskDomain struct {
	taskCountStore storages.TaskCountStore
	taskStore      storages.TaskStore
	userStore      storages.UserStore
}

func (d *taskDomain) GetList(ctx context.Context, createdDate string) ([]*storages.Task, error) {
	userID, ok := utils.ExtractFromContext(ctx)

	if !ok {
		return nil, errors.ErrUserIdIsInvalid
	}

	if _, err := d.userStore.FindUser(ctx, userID); err != nil {
		return nil, errors.ErrUserDoesNotExist
	}

	return d.taskStore.GetTasks(ctx, &storages.Task{
		UserID:      userID,
		CreatedDate: createdDate,
	})
}

func (d *taskDomain) Create(ctx context.Context, content string) (*storages.Task, error) {
	userID, ok := utils.ExtractFromContext(ctx)

	if !ok {
		return nil, errors.ErrUserIdIsInvalid
	}

	user, err := d.userStore.FindUser(ctx, userID)
	if err != nil {
		return nil, errors.ErrUserDoesNotExist
	}

	date := time.Now().Format("2006-01-02")
	total := d.taskCountStore.Inc(utils.GetKey(userID, date))
	// check is amount exceeded
	if total > user.MaxTodo {
		return nil, errors.ErrTaskLimitExceeded
	}

	task := &storages.Task{
		UserID:       userID,
		Content:      content,
		CreatedDate:  date,
		NumberInDate: total,
	}
	if err := d.taskStore.AddTask(ctx, task); err != nil {
		d.taskCountStore.Desc(utils.GetKey(userID, date))
		return nil, err
	}

	return task, nil
}

func NewTaskDomain(
	taskCountStore storages.TaskCountStore,
	taskStore storages.TaskStore,
	userStore storages.UserStore,
) TaskDomain {
	return &taskDomain{
		taskCountStore: taskCountStore,
		taskStore:      taskStore,
		userStore:      userStore,
	}
}
