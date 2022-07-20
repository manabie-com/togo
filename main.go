package main

import (
	"log"
	"net/http"

	"github.com/tonghia/togo/internal/service"
	"github.com/tonghia/togo/internal/store"
)

func main() {

	store, storeErr := store.NewStore()
	if storeErr != nil {
		log.Fatalf("Could not init store Error: %v", storeErr)
	}
	service := service.NewService(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", service.Health)
	mux.HandleFunc("/task", service.RecordTask)

	listenAddr := ":8080"
	err := http.ListenAndServe(listenAddr, mux)
	if err != nil {
		log.Fatalf("Server could not start listening on %s. Error: %v", listenAddr, err)
	}
}
