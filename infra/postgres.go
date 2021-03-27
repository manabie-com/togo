package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewPostgresDB() (*sql.DB, error) {
	var ds = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	db, err := sql.Open(os.Getenv("POSTGRES_DRIVER"), ds)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
