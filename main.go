package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/workers"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	// db, err := sql.Open("sqlite3", "./data.db")
	// Change database to postgres
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=admin password=admin123 dbname=manabie sslmode=disable")

	fmt.Println("Initial Server!")

	if err != nil {
		log.Fatal("error opening db", err)
	}

	defer db.Close()

	// Initial table for database
	db.AutoMigrate(&storages.Task{})
	db.AutoMigrate(&storages.User{})
	db.AutoMigrate(&storages.ConfigServer{})

	// Create user for testing - When testing first time please open this comment code
	/*
	var user = storages.User{
		ID:       "firstUser",
		Password: "example",
	}

	rs := db.Create(&user)

	if rs.Error != nil {
		fmt.Println(rs.Error.Error())
	}
	*/

	// Initial config server for processing max tasks / day of user
	
	var configServers = []storages.ConfigServer{{
		Name:  "config_max_tasks_per_day",
		Value: 5,
	}}

	for index := range configServers {
		var config = configServers[index]
		if db.Where("name=?", config.Name).First(&config).RowsAffected <= 0 {
			db.Create(&config)
		} else {
			db.Model(&storages.ConfigServer{}).Where("name=?", configServers[index].Name).Update("value", configServers[index].Value)
		}
	}

	var configTask storages.ConfigServer
	rsConfig := db.Where("name=?", "config_max_tasks_per_day").First(&configTask)
	if rsConfig.Error != nil {
		fmt.Println(rsConfig.Error.Error())
	}

	/*
		- Function "http.ListenAndServe" will receive Port and a Pointer for Interface "Hanlder".
		- From Interface "Handler" will have a function ServeHTTP(ResponseWriter, *Request).
		-----------------||-----------------
		- Initial ToDoService with some Methods.
		- ToDoService have function "ServeHTTP" => when initial success will run base on interface of Function "http.ListenAndServe".
	*/

	// Go routine a scheduling here for reset config_max_tasks_per_day of all user when new day 00:00
	// TODO
	//
	go workers.CronResetUserMaxTaskEveryDay(db)

	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &postgres.PostgresDB{
			DB: db,
		},
		MaxTasksConfig: configTask.Value,
	})
}
