package taskrepo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"task_service/modules/taskmodel"
)

type TaskStorage interface {
	ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error)
	CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error
	GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error)
	UpdateTask(ctx context.Context, cond map[string]interface{}, dataUpdate *taskmodel.TaskUpdate) error
	DeleteTask(ctx context.Context, cond map[string]interface{}) error
}

type TaskCacheStorage interface {
	IncrBy(ctx context.Context, key string, number int) (int, error)
	Get(ctx context.Context, key string) (string, error)
}

type repo struct {
	store      TaskStorage
	cacheStore TaskCacheStorage
}

func NewRepo(store TaskStorage, cacheStore TaskCacheStorage) *repo {
	return &repo{store: store, cacheStore: cacheStore}
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

func (r *repo) GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error) {
	task, err := r.store.GetTask(ctx, cond)
	if err != nil {
		return nil, sdkcm.ErrCannotGetEntity("task", err)
	}

	return task, nil
}

func (r *repo) UpdateTask(ctx context.Context, cond map[string]interface{}, dataUpdate *taskmodel.TaskUpdate) error {
	if err := r.store.UpdateTask(ctx, cond, dataUpdate); err != nil {
		return sdkcm.ErrCannotUpdateEntity("task", err)
	}

	return nil
}

func (r *repo) DeleteTask(ctx context.Context, cond map[string]interface{}) error {
	if err := r.store.DeleteTask(ctx, cond); err != nil {
		return sdkcm.ErrCannotDeleteEntity("task", err)
	}

	return nil
}

func (r *repo) IncrByNumberTaskToday(ctx context.Context, userId, number int) (int, error) {
	key := fmt.Sprintf("users/limit-tasks/%d/%s", userId, time.Now().Format("01-02-2006"))
	number, err := r.cacheStore.IncrBy(ctx, key, number)
	if err != nil {
		return 0, sdkcm.ErrCannotCreateEntity("task", err)
	}

	return number, nil
}

func (r *repo) CountTaskToday(ctx context.Context, userId int) (int, error) {
	key := fmt.Sprintf("users/limit-tasks/%d/%s", userId, time.Now().Format("01-02-2006"))
	number, err := r.cacheStore.Get(ctx, key)
	if err != nil {
		return 0, sdkcm.ErrEntityNotFound("task", err)
	}

	count, err := strconv.Atoi(number)
	if err != nil {
		return 0, sdkcm.ErrEntityNotFound("task", err)
	}

	return count, nil
}
