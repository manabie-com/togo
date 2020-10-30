package usecase

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
)

type TogoUsecase interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error)
	AddTask(ctx context.Context, t *entities.Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}

type togoUsecase struct {
	togoStorages storages.DB
}

func NewTogoUsecase(d storages.DB) TogoUsecase {

}
