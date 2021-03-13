package interfaces

import (
	"context"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/usecase/response"
)

type Presenter struct {
	User UserPresenter
	Task TaskPresenter
}

type UserPresenter interface {
	SignUp(ctx context.Context, u *entity.User) (*response.SignUp, error)
	SignIn(ctx context.Context, u *entity.User) (*response.SignIn, error)
}

type TaskPresenter interface {
	Fetch(ctx context.Context, ts []*entity.Task) *response.GetTasks
	Create(ctx context.Context, t *entity.Task) *response.CreateTask
}
