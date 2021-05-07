package storages

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type IDatabase interface {
	GetMaxTodo(ctx context.Context, userID string) (int32, error)
	CountTasks(context.Context, string, string) (int32, error)
	RetrieveTasks(context.Context, string, string) ([]*Task, error)
	AddTask(context.Context, *Task, func(string, string) error) error
	ValidateUser(context.Context, string, string) bool
	Close()
}

func NewDatabase(cfg *config.Config) (IDatabase, error) {
	// if environment is D (development/testing) then sqlite will be database
	// otherwise postgres will be database
	var (
		db *sql.DB
		err error
	)
	if cfg.Environment == "D" {
		db, err = sql.Open("sqlite3", cfg.SQLite)
	} else {
		postgres := cfg.Postgres
		db, err = sql.Open("postgres",
			fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
				postgres.Host, postgres.Port, postgres.User, postgres.Password, postgres.DBName, postgres.SSL))
	}
	if err != nil {
		return nil, err
	}
	return &LiteDB{DB: db}, db.Ping()
}
