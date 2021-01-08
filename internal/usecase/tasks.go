package usecase

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/model"
	"time"
)

func (a Usecase) AddTask(ctx context.Context, userID string, t *model.Task) (*model.Task, error) {
	// add named mutex to avoid race condition when calling this method multiple times concurrently
	a.AddTaskLocker.Lock(userID)
	defer a.AddTaskLocker.Unlock(userID)

	// get user
	u, err := a.Store.User().Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	// count user's current tasks
	now := time.Now().UTC().Format("2006-01-02")
	num, err := a.Store.Task().CountTasksByUser(ctx, userID, sql.NullString{now, true})
	if err != nil {
		return nil, err
	}

	// check tasks limit
	if !(num < u.MaxTodo) {
		return nil, model.NewError(model.ErrTasksLimitExceeded, "")
	}

	// persist to db
	t.ID = uuid.New().String()
	res, err := a.Store.Task().AddTask(ctx, userID, t)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a Usecase) ListTasks(ctx context.Context, userID string, createdDate sql.NullString) ([]*model.Task, error) {
	res, err := a.Store.Task().RetrieveTasks(ctx, userID, createdDate)
	if err != nil {
		return nil, err
	}

	return res, nil
}
