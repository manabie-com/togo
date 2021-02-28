package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/db"
	"github.com/manabie-com/togo/internal/handlers"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

var database *sql.DB
var httpHandler *handlers.HttpHandler

func testMain(m *testing.M) int {
	var err error

	database, err = db.SetupPostgres(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"postgres",
		"123456",
		"postgres",
		5432,
		"testdb",
	))
	if err != nil {
		log.Printf("error when connecting to postgres: %s\n", err.Error())
		return 1
	}

	if _, err = database.Exec(db.Schema); err != nil {
		log.Printf("error when applying database schema %s\n", err.Error())
		return 1
	}

	redis, err := db.SetupRedis("redis:6379", "", 0)
	if err != nil {
		log.Printf("error when connecting to redis: %s\n", err.Error())
	}

	httpHandler = &handlers.HttpHandler{
		UserService: &services.UserService{
			JWTKey: "secret",
			Storage: &postgres.PostgresDB{
				DB: database,
			},
		},
		TaskService: &services.TaskService{
			Redis: redis,
			Storage: &postgres.PostgresDB{
				DB: database,
			},
		},
	}

	return m.Run()
}
