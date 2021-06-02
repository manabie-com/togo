package main

import (
	"path/filepath"

	internal "github.com/manabie-com/togo/internal"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dataSourceName, _ := filepath.Abs("./data.db")
	s := internal.NewServer("sqlite3", dataSourceName)
	s.ListenAndServe()
}
