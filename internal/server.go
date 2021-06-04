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

func NewSQLDataBaseFromDriver(driverName, dataSourceName string) *storages.LiteDBAdapter {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	switch driverName {
	case "sqlite3":
		return &storages.LiteDBAdapter{DB: db}
	default:
		log.Fatalf("Not supported SQL Driver:%s", driverName)
		return &storages.LiteDBAdapter{}
	}
}

func NewServer(driverName, dataSourceName, port, jwtSecret string) *http.Server {
	st := NewSQLDataBaseFromDriver(driverName, dataSourceName)

	r := mux.NewRouter()

	as := &services.AuthService{
		JWTSecret: jwtSecret,
		Store:     st,
	}

	tds := &services.ToDoService{
		Store: st,
		Auth:  as,
	}

	r.HandleFunc("/login", as.IssueJWTToken)
	r.HandleFunc("/tasks", tds.ServeHTTP)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	return s
}
