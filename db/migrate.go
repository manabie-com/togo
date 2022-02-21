package db

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
)

func Migrate(sourceURL, dbURL string) error {
	m, err := migrate.New(
		sourceURL,
		dbURL)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		log.Fatal(err)
		return err
	}

	log.Infof("Migrate Done")
	return nil
}
