package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/internal/driver"
)

const portNumber = ":3000"

func run() (*http.Server, *driver.DB, error) {
	// connect to database
	db, err := driver.ConnectDB(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatal("cannot connect to database! dying...")
	}

	{
		// repository pattern
	}

	{
		// seed user
		var id int
		query := "SELECT id FROM users WHERE username = 'khxingn' AND password = 'Qq@1234567'"
		row := db.SQL.QueryRowContext(context.Background(), query)
		err = row.Scan(&id)

		if err != nil {
			insertQuery := "INSERT INTO users(id, username, password) VALUES (1, 'khxingn', 'Qq@1234567');"
			_, err := db.SQL.ExecContext(context.Background(), insertQuery)
			if err != nil {
				return nil, db, err
			}
		}
	}

	log.Println("Starting server...")

	return &http.Server{Addr: portNumber, Handler: routes()}, db, nil
}

func main() {
	srv, db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	if err = srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
