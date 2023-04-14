package taskhandler

import (
	"context"
	"errors"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"task_service/modules/taskmodel"
)

type UpdateTaskRepo interface {
	GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error)
	UpdateTask(ctx context.Context, cond map[string]interface{}, dataUpdate *taskmodel.TaskUpdate) error
}

type updateTaskHdl struct {
	repo      UpdateTaskRepo
	requester *sdkcm.SimpleUser
}

func NewUpdateTaskHdl(repo UpdateTaskRepo, requester *sdkcm.SimpleUser) *updateTaskHdl {
	return &updateTaskHdl{repo: repo, requester: requester}
}

func (h *updateTaskHdl) Response(ctx context.Context, id int, dataUpdate *taskmodel.TaskUpdate) error {
	data, err := h.repo.GetTask(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return sdkcm.ErrCannotGetEntity("task", err)
	}

	isOwner := h.requester.ID == data.UserId
	if !isOwner {
		return sdkcm.ErrNoPermission(errors.New("no permission"))
	}

	if err = h.repo.UpdateTask(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return sdkcm.ErrCannotUpdateEntity("task", err)
	}

	return nil
}
