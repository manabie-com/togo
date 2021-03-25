package postgres

import (
	"context"

	"github.com/manabie-com/togo/internal/storages"
)

type Querier interface {
	AddTask(ctx context.Context, arg *storages.Task) error
	RetrieveTasks(ctx context.Context, arg RetrieveTasksParams) ([]storages.Task, error)
	ValidateUser(ctx context.Context, arg ValidateUserParams) bool
	CountTaskPerDay(ctx context.Context, arg CountTaskPerDayParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
