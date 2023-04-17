package taskhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/common"
	"togo/modules/task/taskmodel"
)

type CreateTaskRepo interface {
	CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error
	CountTaskToday(ctx context.Context, userId int) (int, error)
	IncrByNumberTaskToday(ctx context.Context, userId, number int) (int, error)
}

type CreateTaskUserRepo interface {
	GetUserLimit(ctx context.Context, userId int) (int, error)
}

type createTaskHdl struct {
	repo      CreateTaskRepo
	userRepo  CreateTaskUserRepo
	requester *sdkcm.SimpleUser
}

func NewCreateTaskHdl(repo CreateTaskRepo, userRepo CreateTaskUserRepo, requester *sdkcm.SimpleUser) *createTaskHdl {
	return &createTaskHdl{repo: repo, userRepo: userRepo, requester: requester}
}

func (h *createTaskHdl) Response(ctx context.Context, data *taskmodel.TaskCreate) error {
	data.UserId = h.requester.ID

	limit, err := h.userRepo.GetUserLimit(ctx, h.requester.ID)
	if err != nil {
		return common.ErrCannotGetUserLimit
	}

	numberTaskToday, err := h.repo.IncrByNumberTaskToday(ctx, h.requester.ID, 1)
	if err != nil {
		taskToday, err := h.repo.CountTaskToday(ctx, h.requester.ID)
		if err != nil {
			return sdkcm.ErrCannotCreateEntity("task", err)
		}

		numberTaskToday = taskToday
	}

	if numberTaskToday > limit {
		h.repo.IncrByNumberTaskToday(ctx, h.requester.ID, -1)

		return common.ErrLimitTaskToday
	}

	if err = h.repo.CreateTask(ctx, data); err != nil {
		h.repo.IncrByNumberTaskToday(ctx, h.requester.ID, -1)

		return sdkcm.ErrCannotCreateEntity("task", err)
	}

	return nil
}
