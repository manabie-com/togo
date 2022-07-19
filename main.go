package main

import (
	"io"
	"log"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

func main() {

	listenAddr := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheckHandler)

	err := http.ListenAndServe(listenAddr, mux)
	if err != nil {
		log.Fatalf("Server could not start listening on %s. Error: %v", listenAddr, err)
	}
}
