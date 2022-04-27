package main

import (
	"fmt"
	"net/http"
	"manabie.com/internal/views"
)

func main() {
	fmt.Println("start")
	http.HandleFunc("/tasks", views.TaskViewRest)
	http.ListenAndServe(":8090", nil)
}