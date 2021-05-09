package storages

import (
	"fmt"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IDatabase interface {
	AddUser(userID, password string, maxTodo int32) error
	GetMaxTodo(userID string) (int32, error)
	CountTasks(string, string) (int32, error)
	RetrieveTasks(string, string) ([]*models.Task, error)
	AddTask(*models.Task, func(string, string) error) error
	ValidateUser(string, string) bool
}

func NewDatabase(cfg *config.Config) (IDatabase, error) {
	// if environment is D (development/testing) then sqlite will be chosen
	// otherwise postgres will be chosen
	var (
		db *gorm.DB
		err error
	)
	if cfg.Environment == "D" {
		db, err = gorm.Open(sqlite.Open(cfg.SQLite), &gorm.Config{})
	} else {
		pg := cfg.Postgres
		dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			pg.Host, pg.Port, pg.User, pg.Password, pg.DBName, pg.SSL)
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	}
	if err != nil {
		return nil, err
	}
	// enable debug mode
	db = db.Debug()
	return NewStore(db), nil
}
