package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/tonghia/togo/internal/service"
	"github.com/tonghia/togo/internal/store"
	"github.com/tonghia/togo/pkg/database"
	"github.com/tonghia/togo/pkg/tasklimit"
)

type Config struct {
	ListenAddr  string
	MySQLConfig database.MySQLConfig
}

var config = Config{
	ListenAddr: ":8080",
	MySQLConfig: database.MySQLConfig{
		Host:     "0.0.0.0",
		Database: "godb",
		Port:     3306,
		Username: "gouser",
		Password: "gopassword",
		Options:  "checkConnLiveness=false&loc=Local&parseTime=true&maxAllowedPacket=0",
	},
}

func main() {

	db := database.NewMySQLDatabase(config.MySQLConfig)
	mainStore := store.New(db)

	service := service.NewService(mainStore, tasklimit.GetUserLimiSvc())

	server := mux.NewRouter()
	server.HandleFunc("/health", service.Health)
	server.HandleFunc("/user/{userID}/task", service.RecordTask)

	err := http.ListenAndServe(config.ListenAddr, server)
	if err != nil {
		log.Fatalf("Server could not start listening on %s. Error: %v", config.ListenAddr, err)
	}
}
