package helpers

import (
	"database/sql"
	"log"
	"togo/data"
)

func EnsureTablesExist(db *sql.DB) {
	if _, err := db.Exec(data.UsersTableCreationQuery); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(data.TodosTableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func CreateInitialUser(db *sql.DB) {
	if _, err := db.Exec(data.InitialUserQuery); err != nil {
		log.Fatal(err)
	}
}

func ClearTables(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS todos")
	db.Exec("DROP TABLE IF EXISTS users")
}
