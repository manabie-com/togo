package migrations

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateUp(driver, dbUrl string) bool {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/db",
	}
	db, err := sql.Open(driver, dbUrl)
	if err != nil {
		fmt.Println("Error Connection", err)
		return false
	}

	n, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error migration", err)
		return false
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return true
}

func MigrateDown(driver, dbUrl string) bool {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/db",
	}
	db, err := sql.Open(driver, dbUrl)
	if err != nil {
		fmt.Println("Error Connection", err)
		return false
	}

	n, err := migrate.Exec(db, driver, migrations, migrate.Down)
	if err != nil {
		fmt.Println("Error migration", err)
		return false
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return true
}
