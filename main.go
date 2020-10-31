package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	newLiteDB := &sqlite.LiteDB{DB: db}
	todoUs := usecase.NewTogoUsecase(newLiteDB)
	mux := chi.NewRouter()
	transport.NewTogoHandler(mux, &todoUs, "wqGyEBBfPK9w3Lxw")
	http.ListenAndServe(":5050", mux)
}
