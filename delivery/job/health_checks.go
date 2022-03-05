package job

import (
	"net/http"

	"github.com/robfig/cron/v3"

	"github.com/khangjig/togo/client/logger"
)

type healthChecks struct {
	url string
}

func NewHealthChecks(url string) IJob {
	return healthChecks{
		url: url,
	}
}

func (h healthChecks) callHealthCheck() {
	if h.url != "" {
		resp, err := http.Get(h.url)
		if err != nil {
			logger.GetLogger().Fatal("failed to call request to health check")

			return
		}

		if resp != nil {
			_ = resp.Body.Close()
		}
	}
}

func (h healthChecks) Run() {
	c := cron.New()

	_, err := c.AddFunc("*/1 * * * *", func() { h.callHealthCheck() })
	if err != nil {
		logger.GetLogger().Fatal("failed to schedule health check short period")

		return
	}

	c.Start()
}
