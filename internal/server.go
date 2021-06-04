package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	services "github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	_ "github.com/mattn/go-sqlite3"
)

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

	r.HandleFunc("/login", as.IssueJWTToken)
	r.HandleFunc("/tasks", tds.ServeHTTP)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	return s
}
