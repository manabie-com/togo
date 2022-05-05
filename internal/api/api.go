package api

import (
	"fmt"
	"github.com/jmsemira/togo/internal/middleware"

	"net/http"
)

// function to initialize api and set routes
func InitApi() {
	http.Handle("/create_todo", middleware.AccessTokenMiddleware(http.HandlerFunc(CreateTodoHandler)))
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)

	// start http server
	fmt.Println("Starting API server")
	http.ListenAndServe(":8080", nil)
}
