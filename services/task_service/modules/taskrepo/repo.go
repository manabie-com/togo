package taskrepo

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"task_service/modules/taskmodel"
)

type TaskStorage interface {
	ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error)
	CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error
}

type repo struct {
	store TaskStorage
}

func NewRepo(store TaskStorage) *repo {
	return &repo{store: store}
}

func (r *repo) ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error) {
	tasks, err := r.store.ListItem(ctx, filter, paging)
	if err != nil {
		return nil, sdkcm.ErrCannotListEntity("tasks", err)
	}

	return tasks, nil
}

func (r *repo) CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error {
	if err := r.store.CreateTask(ctx, data); err != nil {
		return err
	}

	return nil
}
