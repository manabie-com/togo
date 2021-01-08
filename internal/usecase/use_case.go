package usecase

import (
	"github.com/manabie-com/togo/internal/lib"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgresql"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

type Usecase struct {
	Store           storages.Store
	AddTaskLocker *lib.NamedMutex
}

func New(conf model.AppSettings) (*Usecase, error) {
	var store storages.Store
	var err error
	switch conf.DatabaseSettings.Provider {
	case model.DatabaseProviderSQLite:
		store, err = sqllite.NewStore(conf.DatabaseSettings)
		if err != nil {
			return nil, err
		}
	case model.DatabaseProviderPostgresql:
		fallthrough
	default:
		store, err = postgresql.NewStore(conf.DatabaseSettings)
		if err != nil {
			return nil, err
		}
	}

	a := Usecase{
		Store: store,
		AddTaskLocker: lib.NewNamedMutex(),
	}

	return &a, nil
}
