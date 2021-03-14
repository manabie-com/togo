package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetConnection() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	maxAttempt, _ := strconv.Atoi(os.Getenv("DB_NUM_ATTEMPT"))
	if err := WaitingForDB(conn, maxAttempt); err != nil {
		return nil, err
	}

	return conn, nil
}
