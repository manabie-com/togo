package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	pg "github.com/manabie-com/togo/internal/storages/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	// Initialize Datasource Connection & Migrates
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("error opening db", err)
	}
	db.AutoMigrate(&storages.User{}, &storages.Task{})

	// Start server
	http.ListenAndServe(":5050", &services.ToDoService{
		JWTKey: os.Getenv("JWT_KEY"),
		Store: &pg.Storage{
			DB: db,
		},
	})
}
