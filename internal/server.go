package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	services "github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func NewServer(driverName, dataSourceName, port, jwtSecret string) *http.Server {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	as := &services.AuthService{JWTSecret: jwtSecret}
	tds := &services.ToDoService{Auth: as}

	switch driverName {
	case "sqlite3":
		as.Store = &storages.LiteDBAdapter{DB: db}
		tds.Store = &storages.LiteDBAdapter{DB: db}
	case "postgres":
		as.Store = &storages.PGDBAdapter{DB: db}
		tds.Store = &storages.PGDBAdapter{DB: db}
	}

	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	r.HandleFunc("/login", as.IssueJWTToken).Methods(http.MethodGet)
	r.HandleFunc("/tasks", tds.ServeHTTP).Methods(http.MethodGet, http.MethodPost)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	return s
}
