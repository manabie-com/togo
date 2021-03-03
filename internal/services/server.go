package services

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"log"
	"net/http"
)

func StartServer(db *sql.DB) {

	s := &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgreDB{
			DB: db,
		},
	}

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(corsMiddleware)
	r.PathPrefix("/").Methods(http.MethodOptions).HandlerFunc(methodOptionsHandler)
	r.HandleFunc("/login", s.loginHandler)

	taskR := r.PathPrefix("/task").Subrouter()
	taskR.Use(s.authenticateMiddleware)
	taskR.Methods(http.MethodGet).HandlerFunc(s.getTasksHandler)
	taskR.Methods(http.MethodPost).HandlerFunc(s.addTaskHandler)

	err := http.ListenAndServe(":5050", r)
	log.Fatal(err)
}