package taskhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"task_service/common"
	"task_service/modules/taskmodel"
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
	revertNumber := 0
	defer h.repo.IncrByNumberTaskToday(ctx, h.requester.ID, revertNumber)

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
		revertNumber = -1

		return common.ErrLimitTaskToday
	}

	if err = h.repo.CreateTask(ctx, data); err != nil {
		revertNumber = -1

		return sdkcm.ErrCannotCreateEntity("task", err)
	}

	return nil
}
