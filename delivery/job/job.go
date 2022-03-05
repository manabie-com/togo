package job

import (
	"github.com/khangjig/togo/config"
	"github.com/khangjig/togo/repository"
)

type IJob interface {
	Run()
}

type Jobs []IJob

func (js Jobs) Run() {
	for _, j := range js {
		go j.Run()
	}
}

func New(repo *repository.Repository) Jobs {
	return Jobs{
		NewHealthChecks(config.GetConfig().HealthCheck.EndPoint),
		NewResetMaxDailyLimit(repo),
	}
}
