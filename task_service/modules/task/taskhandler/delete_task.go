package taskhandler

import (
	"context"
	"errors"

	"github.com/golang-module/carbon/v2"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/task/taskmodel"
)

type DeleteTaskRepo interface {
	GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error)
	DeleteTask(ctx context.Context, cond map[string]interface{}) error
	IncrByNumberTaskToday(ctx context.Context, userId, number int) (int, error)
}

type deleteTaskHdl struct {
	repo      DeleteTaskRepo
	requester *sdkcm.SimpleUser
}

func NewDeleteTaskHdl(repo DeleteTaskRepo, requester *sdkcm.SimpleUser) *deleteTaskHdl {
	return &deleteTaskHdl{repo: repo, requester: requester}
}

func (h *deleteTaskHdl) Response(ctx context.Context, id int) error {
	data, err := h.repo.GetTask(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return sdkcm.ErrCannotGetEntity("task", err)
	}

	isOwner := h.requester.ID == data.UserId
	if !isOwner {
		return sdkcm.ErrNoPermission(errors.New("no permission"))
	}

	if err = h.repo.DeleteTask(ctx, map[string]interface{}{"id": id}); err != nil {
		return sdkcm.ErrCannotDeleteEntity("task", err)
	}

	if data.CreatedDate.IsSameDay(carbon.Now()) {
		h.repo.IncrByNumberTaskToday(ctx, h.requester.ID, -1)
	}

	return nil
}
