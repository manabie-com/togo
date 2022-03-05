package job

import (
	"context"

	"github.com/robfig/cron/v3"

	"github.com/khangjig/togo/client/logger"
	"github.com/khangjig/togo/repository"
)

type resetMaxDailyLimit struct {
	repo *repository.Repository
}

func NewResetMaxDailyLimit(repo *repository.Repository) IJob {
	return resetMaxDailyLimit{
		repo: repo,
	}
}

func (r resetMaxDailyLimit) callResetMaxDailyLimit() {
	err := r.repo.UserCache.ResetTotalTodo(context.Background())
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}
}

func (r resetMaxDailyLimit) Run() {
	c := cron.New()

	// run every mid night
	_, err := c.AddFunc("0 0 * * *", func() { r.callResetMaxDailyLimit() })
	if err != nil {
		logger.GetLogger().Fatal("failed to schedule reset maximum daily limit")

		return
	}

	c.Start()
}
