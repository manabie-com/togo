package services

import (
	"context"
	"github.com/manabie-com/togo/internal/storages/repos"
	"github.com/manabie-com/togo/internal/tools"
	"time"
)

type IQuotaService interface {
	LimitTask(ctx context.Context) error
}

type QuotaService struct {
	repo repos.IQuotaRepo
}

func (qs *QuotaService) LimitTask(ctx context.Context) error {
	now := time.Now()
	userID, err := tools.UserIDFromCtx(ctx)
	if err != nil {
		return err
	}
	dateNow := now.Format("2006-01-02")
	countTaskPerDate := qs.repo.CountTaskPerDay(ctx, userID, dateNow)
	maxTaskPerDate := qs.repo.GetLimitPerUser(ctx, userID)
	if countTaskPerDate >= maxTaskPerDate {
		return tools.NewTodoError(405, "You reach a limit to create task in date")
	}
	return nil
}

func NewQuotaService(repo repos.IQuotaRepo) IQuotaService {
	return &QuotaService{
		repo: repo,
	}
}
