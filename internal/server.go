package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	services "github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
)

func NewServer(driverName, dataSourceName, port, jwtSecret string) *http.Server {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	r := mux.NewRouter()

	as := &services.AuthService{
		JWTSecret: jwtSecret,
		Store: &storages.LiteDBAdapter{
			DB: db,
		},
	}

	tds := &services.ToDoService{
		JWTKey: jwtSecret,
		Store: &storages.LiteDBAdapter{
			DB: db,
		},
		Auth: as,
	}

	r.HandleFunc("/login", as.IssueJWTToken)
	r.HandleFunc("/tasks", tds.ServeHTTP)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	return s
}
