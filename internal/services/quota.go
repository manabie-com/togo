package services

import (
	"context"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
	"time"
)

type IQuotaService interface {
	LimitTask(ctx context.Context) *tools.TodoError
}

type QuotaService struct {
	repo        storages.IQuotaRepo
	contextTool tools.IContextTool
}

func (qs *QuotaService) LimitTask(ctx context.Context) *tools.TodoError {
	now := time.Now()
	userID, err := qs.contextTool.UserIDFromCtx(ctx)
	if err != nil {
		return err
	}
	dateNow := now.Format("2006-01-02")
	countTaskPerDate, err := qs.repo.CountTaskPerDayStore(ctx, storages.CountTaskPerDayParams{UserID: userID, CreatedDate: dateNow})
	if err != nil {
		return err
	}
	maxTaskPerDate, err := qs.repo.GetLimitPerUserStore(ctx, userID)
	if err != nil {
		return err
	}
	if countTaskPerDate >= int64(maxTaskPerDate) {
		return tools.NewTodoError(http.StatusMethodNotAllowed, "You reach a limit to create task in date")
	}
	return nil
}

func NewQuotaService(repo storages.IQuotaRepo, contextTool tools.IContextTool) IQuotaService {
	return &QuotaService{
		repo:        repo,
		contextTool: contextTool,
	}
}
