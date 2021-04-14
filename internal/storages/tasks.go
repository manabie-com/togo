package storages

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/services/tasks"

	"github.com/go-pg/pg"

	"github.com/google/uuid"
)

type task struct {
	tableName struct{} `sql:"tasks" pg:",discard_unknown_columns"`

	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type taskRepo struct {
	db *pg.DB
}

func NewTaskRepo(db *pg.DB) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

func (r *taskRepo) Save(ctx context.Context, t *tasks.Task) error {
	return r.db.WithContext(ctx).Insert(&task{
		ID:        uuid.New(),
		UserID:    t.UserID,
		Content:   t.Content,
		CreatedAt: t.CreatedAt,
	})
}
