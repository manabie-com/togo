package postgresql

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

type Store struct {
	db        sqlbuilder.Database
	userStore storages.User
	taskStore storages.Task
}

func NewStore(conf model.DatabaseSettings) (*Store, error) {
	dbx, err := sql.Open("postgres", conf.ConnectionString)
	if err != nil {
		return nil, err
	}

	db, err := postgresql.New(dbx)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(32)

	// try first query
	_, err = db.Exec("select 1")
	if err != nil {
		return nil, err
	}

	return &Store{
		db:        db,
		userStore: NewUserStore(db),
		taskStore: NewTaskStore(db),
	}, nil
}

func (s Store) User() storages.User {
	return s.userStore
}

func (s Store) Task() storages.Task {
	return s.taskStore
}

func (s Store) DropAllRecords() {
	s.db.Exec("truncate table " + usersTable + " cascade")
	s.db.Exec("truncate table " + tasksTable + " cascade")
}