package domain

import (
	"context"
	"time"

	"manabie/togo/common/errors"
	"manabie/togo/internal/model"
	"manabie/togo/utils"
)

type TaskDomain interface {
	GetList(ctx context.Context, createdDate string) ([]*model.Task, error)
	Create(ctx context.Context, content string) (*model.Task, error)
}

type taskDomain struct {
	taskCountStore model.TaskCountStore
	taskStore      model.TaskStore
	userStore      model.UserStore
}

func (d *taskDomain) GetList(ctx context.Context, createdDate string) ([]*model.Task, error) {
	userID, ok := utils.ExtractFromContext(ctx)

	if !ok {
		return nil, errors.ErrUserIDIsInvalid
	}

	if _, err := d.userStore.FindUser(ctx, userID); err != nil {
		return nil, errors.ErrUserDoesNotExist
	}

	return d.taskStore.RetrieveTasks(ctx, &model.Task{
		UserID:      userID,
		CreatedDate: createdDate,
	})
}

func (d *taskDomain) Create(ctx context.Context, content string) (*model.Task, error) {
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
	total, err := d.taskCountStore.Inc(ctx, userID, date)
	if err != nil {
		return nil, err
	}
	// check is amount exceeded
	if total > user.MaxTodo {
		return nil, errors.ErrTaskLimitExceeded
	}

	task := &model.Task{
		UserID:       userID,
		Content:      content,
		CreatedDate:  date,
		NumberInDate: total,
	}
	if err := d.taskStore.AddTask(ctx, task); err != nil {
		d.taskCountStore.Desc(ctx, userID, date)
		return nil, err
	}
	return task, nil
}

func NewTaskDomain(
	taskCountStore model.TaskCountStore,
	taskStore model.TaskStore,
	userStore model.UserStore,
) TaskDomain {
	return &taskDomain{
		taskCountStore: taskCountStore,
		taskStore:      taskStore,
		userStore:      userStore,
	}
}
