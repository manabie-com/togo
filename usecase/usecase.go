package usecase

import (
	"github.com/go-kit/kit/log"

	"github.com/valonekowd/togo/usecase/interactor"
	"github.com/valonekowd/togo/usecase/interfaces"
)

type Usecase struct {
	User interactor.UserInteractor
	Task interactor.TaskInteractor
}

func New(ds interfaces.DataSource, presenter interfaces.Presenter, logger log.Logger) Usecase {
	return Usecase{
		User: interactor.NewUserInteractor(ds, presenter.User, logger),
		Task: interactor.NewTaskInteractor(ds, presenter.Task, logger),
	}
}
