package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

func NewServer(driverName, dataSourceName string) *http.Server {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	s := &http.Server{
		Addr: ":5050",
		Handler: &services.ToDoService{
			JWTKey: "wqGyEBBfPK9w3Lxw",
			Store: &sqllite.LiteDB{
				DB: db,
			},
		},
	}

	return s
}
