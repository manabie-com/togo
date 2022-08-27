package main

import (
	task "backend_test/task"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	taskRouter := task.PaymentHandler(r)

	r.Handle("/", taskRouter)

	log.Fatal(http.ListenAndServe(":3000", r))
}
