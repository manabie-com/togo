package sqllite

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"log"
)

// LiteDB for working with sqllite
type LiteDB struct {
	DB        *sql.DB
	userStore storages.User
	taskStore storages.Task
}

func NewStore(conf model.DatabaseSettings) (*LiteDB, error) {
	db, err := sql.Open(conf.SQLite.DriverName, conf.SQLite.DataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	return &LiteDB{
		DB:        db,
		userStore: NewUserStore(db),
		taskStore: NewTaskStore(db),
	}, nil
}

func (s LiteDB) User() storages.User {
	return s.userStore
}

func (s LiteDB) Task() storages.Task {
	return s.taskStore
}

func (s LiteDB) DropAllRecords() {}
