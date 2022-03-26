package postgres

import (
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

func Migrate(dbURL string) error {
	driver, err := iofs.New(fs, "migrations")
	if err != nil {
		return fmt.Errorf("new iofs migration: %w", err)
	}

	migration, err := migrate.NewWithSourceInstance("iofs", driver, dbURL)
	if err != nil {
		return fmt.Errorf("new migrate: %w", err)
	}
	if err := migration.Up(); err != nil {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}
