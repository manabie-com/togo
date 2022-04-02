package migration

import (
	"flag"
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var logger = log.New("migration")

var migrateDown = flag.Bool("down", false, "Undo the last migration or undo til the specific --version")
var migrateVersion = flag.String("version", "", "Exec the migrations up/down to the given migration that matches")

// DefaultMigrationOptions contains default options for the gormigrate package
var DefaultMigrationOptions = &gormigrate.Options{
	TableName:      "migrations",
	IDColumnName:   "id",
	IDColumnSize:   255,
	UseTransaction: true,
}

// Run executes the migrations given
func Run(db *gorm.DB, migrations []*gormigrate.Migration) {
	logger.SetHeader("${time_rfc3339_nano} - [${level}]")
	parseFlags()

	m := gormigrate.New(db, DefaultMigrationOptions, migrations)

	var err error
	if *migrateDown {
		if *migrateVersion == "" {
			err = m.RollbackLast()
		} else {
			err = m.RollbackTo(*migrateVersion)
		}
	} else {
		if *migrateVersion == "" {
			err = m.Migrate()
		} else {
			err = m.MigrateTo(*migrateVersion)
		}
	}

	if err != nil {
		logger.Fatalf("Migration failed: %v", err)
	}

	logger.Info("Migration completed")
}

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: go run cmd/migration/main.go [--down] [--version 200601021504]\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}
