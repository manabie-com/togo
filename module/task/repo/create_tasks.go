package repo

import (
	"context"
	"errors"
	model2 "togo/module/task/model"
	"togo/module/userconfig/model"
	model3 "togo/module/usertask/model"
)

type CreateTaskStore interface {
	CreateTasks(ctx context.Context, data []model2.CreateTask) error
}

type GetUserCfgStore interface {
	Get(ctx context.Context, cond map[string]interface{}) (*model.UserConfig, error)
}

type CreateUserTaskStore interface {
	CreateUserTasks(ctx context.Context, data []model3.CreateUserTask) error
}

type createTaskRepo struct {
	userStore GetUserCfgStore
	taskStore CreateTaskStore
	createUserTaskStore CreateUserTaskStore
}

func NewCreateTaskRepo(userStore GetUserCfgStore, taskStore CreateTaskStore, createUserTaskStore CreateUserTaskStore) *createTaskRepo {
	return &createTaskRepo{userStore: userStore, taskStore: taskStore, createUserTaskStore: createUserTaskStore}
}

func (u *createTaskRepo) CreateTasks(ctx context.Context, userId uint, data []model2.CreateTask) error {
	cond := map[string]interface{}{
		"user_id": userId,
	}

	cfg, err := u.userStore.Get(ctx, cond)
	if err != nil {
		return err
	}

	if cfg != nil {
		if int(cfg.MaxTask) < len(data) {
			return errors.New("Length of task is bigger than user's limit task")
		}
	}

	if err := u.taskStore.CreateTasks(ctx, data); err != nil {
		return err
	}

	userTasks := make([]model3.CreateUserTask, 0)

	for _, task := range data {

		userTask := model3.CreateUserTask{
			UserId: &userId,
			TaskId: task.Id,
		}

		userTasks = append(userTasks, userTask)
	}

	if err := u.createUserTaskStore.CreateUserTasks(ctx, userTasks); err != nil {
		return err
	}

	return nil
}