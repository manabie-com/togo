package main

import (
	"fmt"
	"lntvan166/togo/db"
	"lntvan166/togo/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db.Connect()

	routes.HandleRequest()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server started!")

}
