package formatter

import (
	"github.com/valonekowd/togo/infrastructure/auth"
	"github.com/valonekowd/togo/usecase/interfaces"
)

func New(authCfg auth.Config) interfaces.Presenter {
	return interfaces.Presenter{
		User: NewUserFormatter(authCfg),
		Task: NewTaskFormatter(),
	}
}
