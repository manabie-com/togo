package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tonghia/togo/internal/service"
	"github.com/tonghia/togo/internal/store"
	"github.com/tonghia/togo/pkg/database"
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
	},
}

func main() {

	db := database.NewMySQLDatabase(config.MySQLConfig)
	mainStore := store.New(db)

	service := service.NewService(mainStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", service.Health)
	mux.HandleFunc("/task", service.RecordTask)

	err := http.ListenAndServe(config.ListenAddr, mux)
	if err != nil {
		log.Fatalf("Server could not start listening on %s. Error: %v", config.ListenAddr, err)
	}
}
