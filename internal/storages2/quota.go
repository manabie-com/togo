package storages2

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type IQuotaRepo interface {
	CountTaskPerDay(ctx context.Context, arg CountTaskPerDayParams) (int64, error)
	GetLimitPerUser(ctx context.Context, id string) (int32, error)
}

type QuotaRepo struct {
	*Queries
	db *sqlx.DB
}

func NewQuotaRepo(db *sqlx.DB) IQuotaRepo {
	return &QuotaRepo{
		db: db,
	}
}
