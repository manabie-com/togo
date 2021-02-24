package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/manabie-com/togo/pkg/common/crypto"
	"log"

	_ "github.com/lib/pq"
)

func setup() *Store {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "test")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (" +
		"id TEXT NOT NULL PRIMARY KEY, " +
		"password TEXT NOT NULL, " +
		"max_todo INT8 DEFAULT 5 NOT NULL);")
	if err != nil {
		log.Fatal("error creating users table", err)
	}

	hashedPassword, err := crypto.HashPassword("example")
	if err != nil {
		log.Fatal("error hashing password", err)
	}

	_, err = db.Exec("INSERT INTO users (id, password, max_todo) VALUES($1, $2, $3)", "00001", hashedPassword, 5)
	if err != nil {
		log.Fatal("error generating data", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (" +
		"id TEXT NOT NULL PRIMARY KEY, " +
		"content TEXT NOT NULL, " +
		"user_id TEXT NOT NULL REFERENCES users(id)," +
		"created_date TEXT NOT NULL);")
	if err != nil {
		log.Fatal("error creating tasks table", err)
	}

	return &Store{DB: db}
}

func teardown(store *Store) {
	db := store.DB
	defer db.Close()
	_, err := db.Exec("DROP TABLE tasks;")
	if err != nil {
		log.Fatal("error dropping tasks table", err)
	}

	_, err = db.Exec("DROP TABLE users;")
	if err != nil {
		log.Fatal("error dropping users table", err)
	}

}
