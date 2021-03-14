package postgres

import (
	"database/sql"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

func GetMigrateInstance(migratePath string) (*migrate.Migrate, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file:///"+migratePath, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func WaitingForDB(db *sql.DB, numAttempt int) (err error) {
	for i := 0; i < numAttempt; i++ {
		err = db.Ping()

		if err == nil {
			return
		}

		log.Infof("Trying to connect to DB %d time", i+1)
		time.Sleep(1 * time.Second)
	}

	return err
}
