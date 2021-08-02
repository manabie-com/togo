package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/database"
	"github.com/manabie-com/togo/internal/router"

	"github.com/rs/cors"
)

func main() {
	database.SyncDB(false)
	config := config.GetConfig()
	r := router.GetRouter()
	//set global timezone to GMT+7
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	time.Local = loc //
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf(`
	-----------------------------------------------------
	App name: %s
	Version: %s
	Listening Port: %v
	Environment: %s
	-----------------------------------------------------
	`, config.AppName, config.AppVersion, port, config.Environment)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))

}
