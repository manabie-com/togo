package taskhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"task_service/modules/taskmodel"
)

type ListTaskRepo interface {
	ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error)
}

type listTaskHdl struct {
	repo ListTaskRepo
}

func NewListTaskHdl(repo ListTaskRepo) *listTaskHdl {
	return &listTaskHdl{repo: repo}
}

func (h *listTaskHdl) Response(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error) {
	data, err := h.repo.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, sdkcm.ErrCannotListEntity("tasks", err)
	}

	return data, nil
}
