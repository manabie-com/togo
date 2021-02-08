package dbinterface

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type DBInterface interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, *config.ErrorInfo)
	AddTask(ctx context.Context, t *storages.Task) *config.ErrorInfo
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}

func NewDB(driverName string, dbInfo string) (*sql.DB, error) {
	if driverName == config.DBType.Postgres || driverName == config.DBType.Sqlite {
		return sql.Open(driverName, dbInfo)
	}
	return nil, errors.New("driverName is invalid")
}