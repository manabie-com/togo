package job

import "github.com/khangjig/togo/config"

type IJob interface {
	Run()
}

type Jobs []IJob

func (js Jobs) Run() {
	for _, j := range js {
		go j.Run()
	}
}

func New() Jobs {
	return Jobs{
		NewHealthChecks(config.GetConfig().HealthCheck.EndPoint),
	}
}
