package repo

import (
	"context"
	"gorm.io/gorm"
	"togo/module/task/model"
)

type GetTaskStore interface {
	Get(ctx context.Context, cond map[string]interface{}) (*model.Task, error)
}

type getTaskRepo struct {
	store GetTaskStore
}

func NewGetTaskRepo(store GetTaskStore) *getTaskRepo {
	return &getTaskRepo{store: store}
}

func (u *getTaskRepo) GetTask(ctx context.Context, cond map[string]interface{}) (*model.Task, error) {
	usr, err := u.store.Get(ctx, cond)
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return usr, nil
}
