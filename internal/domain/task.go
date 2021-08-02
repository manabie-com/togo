package domain

import (
	"context"
	"time"

	"github.com/manabie-com/togo/common/errors"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils"
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
		return nil, errors.ErrUserIDIsInvalid
	}

	if _, err := d.userStore.FindUser(ctx, userID); err != nil {
		return nil, errors.ErrUserDoesNotExist
	}

	return d.taskStore.RetrieveTasks(ctx, &storages.Task{
		UserID:      userID,
		CreatedDate: createdDate,
	})
}

func (d *taskDomain) Create(ctx context.Context, content string) (*storages.Task, error) {
	userID, ok := utils.ExtractFromContext(ctx)

	if !ok {
		return nil, errors.ErrUserIDIsInvalid
	}
	user, err := d.userStore.FindUser(ctx, userID)
	if err != nil {
		return nil, errors.ErrUserDoesNotExist
	}

	t := time.Now()
	date := t.Format("2006-01-02")
	if err := d.taskCountStore.CreateIfNotExists(ctx, userID, date); err != nil {
		return nil, err
	}
	total, err := d.taskCountStore.UpdateAndGet(ctx, userID, date)
	if err != nil {
		return nil, err
	}
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
