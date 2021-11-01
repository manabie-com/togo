package db

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/quochungphp/go-test-assignment/src/infrastructure/pg_driver"
	"github.com/quochungphp/go-test-assignment/src/pkgs/settings"
)

func Init() (pgSession *pg.DB) {
	// Init postgresql
	pgSession, err := pg_driver.Setup(pg_driver.DBConfiguration{
		Driver:   os.Getenv(settings.DbDriver),
		Host:     os.Getenv(settings.PgHost),
		Port:     os.Getenv(settings.PgPort),
		Database: os.Getenv(settings.PgDB),
		User:     os.Getenv(settings.PgUser),
		Password: os.Getenv(settings.PgPass),
	})
	if err != nil {
		panic(fmt.Sprint("Error while setup postgres driver: ", err))
	}
	return pgSession
}
