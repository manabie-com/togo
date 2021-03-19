package controllers

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func AA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
