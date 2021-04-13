package usecase

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/storages"
)

type uc struct {
	task storages.Task
}

/*func NewUc(db string) UseCase {
	switch db {
	case "postgres":
		return &uc{
			task: storages.NewPgDB(),
		}
	case "litedb":
		return &uc{
			task: storages.NewLiteDB(),
		}
	}
	return nil
}*/

func NewUc(db string) UseCase {
	return &uc{
		task: storages.NewLiteDB(),
	}
}

type UseCase interface {
	Validate(ctx context.Context, user, password sql.NullString) bool
	CreateToken(id, jwtKey string) (string, error)
	ValidToken(token, JWTKey string) (string, bool)
	List(ctx context.Context, id, createdAt string) ([]*entities.Task, error)
	Add(ctx context.Context, id, date string, task *entities.Task) error
}
