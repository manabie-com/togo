package pg_driver

import (
	"context"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// DBConfiguration ...
type DBConfiguration struct {
	Driver        string
	Host          string
	Port          string
	Database      string
	User          string
	Password      string
	RuntimeParams map[string]interface{}
}

type DB struct {
	DB *pg.DB
}

// Setup ...
func Setup(config DBConfiguration) (pgSession *pg.DB, Error error) {
	pgSession = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})
	ctx := context.Background()
	if err := pgSession.Ping(ctx); err != nil {
		return nil, errors.Wrap(err, "Error while checking database connection")
	}
	log.Println("Database connection successful")

	return pgSession, nil
}

// Disconnect ...
func (s *DB) Disconnect() {
	s.DB.Close()
}
