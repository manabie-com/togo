package main

import (
	"fmt"
	"lntvan166/togo/db"
	"lntvan166/togo/routes"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db.Connect()

	route := mux.NewRouter()
	routes.HandleRequest(route)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started!")

}
