package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/app/config"
)

var configs = config.LoadConfigs()

type DBInterface interface {
	Connect() *sql.DB
}

func Connect() *sql.DB {
	URL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Database.Host,
		configs.Database.Port,
		configs.Database.User,
		configs.Database.Pass,
		configs.Database.Name,
		"disable")

	var db *sql.DB
	db, err := sql.Open("postgres", URL)
	log.Println(URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!")
}
