package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/banhquocdanh/togo/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	fmt.Println("Start service on port :5050")

	http.ListenAndServe(":5050", services.NewToDoService(
		"wqGyEBBfPK9w3Lxw",
		services.WithSqlLiteStore(db)),
	)

}
