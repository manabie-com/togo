package storages

import (
	"context"
	"github.com/manabie-com/togo/internal/tools"
	"net/http"
)

type IQuotaRepo interface {
	CountTaskPerDayStore(ctx context.Context, arg CountTaskPerDayParams) (int64, *tools.TodoError)
	GetLimitPerUserStore(ctx context.Context, id string) (int32, *tools.TodoError)
}

type QuotaRepo struct {
	*Queries
}

func (qr *QuotaRepo) CountTaskPerDayStore(ctx context.Context, arg CountTaskPerDayParams) (int64, *tools.TodoError) {
	count, err := qr.CountTaskPerDay(ctx, arg)
	if err != nil {
		return 0, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return count, nil
}

func (qr *QuotaRepo) GetLimitPerUserStore(ctx context.Context, id string) (int32, *tools.TodoError) {
	count, err := qr.GetLimitPerUser(ctx, id)
	if err != nil {
		return 0, tools.NewTodoError(http.StatusInternalServerError, err.Error())
	}
	return count, nil
}

func NewQuotaRepo(db DBTX) IQuotaRepo {
	return &QuotaRepo{
		Queries: New(db),
	}
}
