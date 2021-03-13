package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/usecase/interfaces"
	"github.com/valonekowd/togo/usecase/request"
	"github.com/valonekowd/togo/usecase/response"
	"github.com/valonekowd/togo/util/helper"
)

type TaskInteractor interface {
	Fetch(context.Context, request.GetTasks) (*response.GetTasks, error)
	Create(context.Context, request.CreateTask) (*response.CreateTask, error)
}

func NewTaskInteractor(ds interfaces.DataSource, presenter interfaces.TaskPresenter, logger log.Logger) TaskInteractor {
	var i TaskInteractor
	{
		i = NewBasicTaskInteractor(ds, presenter)
		// i = LoggingInterceptor(logger)(i)
	}
	return i
}

type basicTaskInteractor struct {
	ds        interfaces.DataSource
	presenter interfaces.TaskPresenter
}

var _ TaskInteractor = basicTaskInteractor{}

func NewBasicTaskInteractor(ds interfaces.DataSource, presenter interfaces.TaskPresenter) TaskInteractor {
	return basicTaskInteractor{
		ds:        ds,
		presenter: presenter,
	}
}

func (i basicTaskInteractor) Fetch(ctx context.Context, req request.GetTasks) (*response.GetTasks, error) {
	userID, _ := helper.UserIDFromCtx(ctx)

	ts, err := i.ds.Task.List(ctx, userID, req.CreatedDate)
	if err != nil {
		return nil, fmt.Errorf("fetching tasks: %w", err)
	}

	return i.presenter.Fetch(ctx, ts), nil
}

func (i basicTaskInteractor) Create(ctx context.Context, req request.CreateTask) (*response.CreateTask, error) {
	taskID, err := helper.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("creating task: %w", err)
	}

	userID, _ := helper.UserIDFromCtx(ctx)

	t := &entity.Task{
		ID:        taskID,
		Content:   req.Content,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
	}

	if err := i.ds.Task.Add(ctx, t); err != nil {
		return nil, fmt.Errorf("creating task: %w", err)
	}

	return i.presenter.Create(ctx, t), nil
}
