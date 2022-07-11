package usecase

import (
	"context"
	"database/sql"

	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/datshiro/togo-manabie/internal/interfaces/service/cache"
	postgres_task "github.com/datshiro/togo-manabie/internal/interfaces/task/repository/postgres"
)

type taskUseCase struct {
	taskRepo     domain.TaskRepository
	CacheService cache.CacheService
	DB           *sql.DB
}

func NewTaskUseCase(dbc *sql.DB, cacheService cache.CacheService) domain.TaskUseCase {
	return &taskUseCase{
		DB:           dbc,
		CacheService: cacheService,
		taskRepo:     postgres_task.NewTaskRepository(),
	}
}

func (t *taskUseCase) CreateTask(ctx context.Context, m *models.Task) error {
	return t.taskRepo.CreateOne(ctx, t.DB, m)
}
