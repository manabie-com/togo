package main

import (
	"os"

	internal "github.com/manabie-com/togo/internal"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := internal.NewServer(
		os.Getenv("DATABASE_DRIVER"),
		os.Getenv("DATABASE_SOURCE"),
		os.Getenv("PORT"),
		os.Getenv("JWT_SECRET"),
	)
	s.ListenAndServe()
}
