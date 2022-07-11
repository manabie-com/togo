package domain

import (
	"context"

	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskRepository interface {
	CreateOne(context.Context, boil.ContextExecutor, *models.Task) error
}

type TaskUseCase interface {
	CreateTask(context.Context, *models.Task) error
}

const (
	TaskPriorityLow = iota
	TaskPriorityMedium
	TaskPriorityHigh
)
