package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	services "github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

func NewServer(driverName, dataSourceName, port, jwtSecret string) *http.Server {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	s := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: &services.ToDoService{
			JWTKey: jwtSecret,
			Store: &sqllite.LiteDB{
				DB: db,
			},
		},
	}

	return s
}
