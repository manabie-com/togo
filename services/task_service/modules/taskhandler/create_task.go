package taskhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"task_service/modules/taskmodel"
)

type CreateTaskRepo interface {
	CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error
}

type createTaskHdl struct {
	repo CreateTaskRepo
}

func NewCreateTaskHdl(repo CreateTaskRepo) *createTaskHdl {
	return &createTaskHdl{repo: repo}
}

func (h *createTaskHdl) Response(ctx context.Context, data *taskmodel.TaskCreate) error {
	if err := h.repo.CreateTask(ctx, data); err != nil {
		return sdkcm.ErrCannotCreateEntity("task", err)
	}

	return nil
}
