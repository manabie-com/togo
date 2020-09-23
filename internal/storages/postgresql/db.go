package postgresql

import (
	"context"
	"database/sql"
	"github.com/go-pg/pg"
	"github.com/manabie-com/togo/internal"
	"github.com/manabie-com/togo/internal/storages"
)

// PostgreDB for working with postgreSQL
type PostgreDB struct {
	Conn *pg.DB
}

func (*PostgreDB) FindUserByID(ctx context.Context, userID string) (*storages.User, error) {
	panic("implement me")
}

func (*PostgreDB) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error) {
	panic("implement me")
}

func (*PostgreDB) AddTask(ctx context.Context, t *storages.Task) error {
	panic("implement me")
}

func (*PostgreDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	panic("implement me")
}

func NewTodoRepository(Conn *pg.DB) internal.Repository {
	return &PostgreDB{Conn}
}
