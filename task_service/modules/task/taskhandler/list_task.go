package taskhandler

import (
	"context"

	"togo/modules/task/taskmodel"

	"github.com/phathdt/libs/go-sdk/sdkcm"
)

type ListTaskRepo interface {
	ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error)
}

type listTaskHdl struct {
	repo      ListTaskRepo
	requester *sdkcm.SimpleUser
}

func NewListTaskHdl(repo ListTaskRepo, requester *sdkcm.SimpleUser) *listTaskHdl {
	return &listTaskHdl{repo: repo, requester: requester}
}

func (h *listTaskHdl) Response(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error) {
	filter.UserId = h.requester.ID

	data, err := h.repo.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, sdkcm.ErrCannotListEntity("tasks", err)
	}

	return data, nil
}
