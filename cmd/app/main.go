package main

import (
	"github.com/xrexonx/togo/cmd/app/config/database"
	"github.com/xrexonx/togo/cmd/app/config/environment"
	"github.com/xrexonx/togo/cmd/app/config/migration"
	"github.com/xrexonx/togo/cmd/app/config/routes"
	"log"
	"net/http"
	"time"
)

func main() {

	// Load environment variables
	environment.LoadEnv()

	// Create Database Connection
	database.Init()

	// Initialize Database and tables
	migration.Run()

	// Setup routes handlers
	handler := routes.Init()

	// Start server
	serverHost := environment.GetValue("HOST")
	serverPort := environment.GetValue("PORT")
	_host := serverHost + ":" + serverPort
	server := &http.Server{
		Handler: handler,
		Addr:    _host,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server started on: " + _host)
	log.Fatal(server.ListenAndServe())
}
