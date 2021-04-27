package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/app/config"
	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/app/routes"
	"github.com/manabie-com/togo/internal/services/transport"
)

func Run() {

	apiPort := config.LoadConfigs().App.Port

	db := models.Connect()
	if db != nil {
		defer db.Close()
	}

	models.TestConnection()
	err := models.MigrationDB()
	if err == nil {
		log.Println("Database migration successfully!")
	}

	transportServices := transport.NewTransport()

	appRoutes := routes.NewAppRoutes(transportServices)
	router := mux.NewRouter().StrictSlash(true)

	routes.Install(router, appRoutes)

	fmt.Printf("Api running on port %s\n", apiPort)

	server := &http.Server{Addr: ":" + apiPort, Handler: router}

	//++++START: GRACEFUL SHUTDOWN++++++
	timeWait := 15 * time.Second
	signChan := make(chan os.Signal, 1)

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	//setup a channel listen signal from OS
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Println("Shutting down")
	//setup interval(timeout) listen to stop program and close all connection
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func() {
		log.Println("Close another connection")
		cancel()
	}()
	log.Println("Stop http server")
	if err := server.Shutdown(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
	log.Printf("Completed")
	//++++END: GRACEFUL SHUTDOWN++++++
}
