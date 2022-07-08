package main

import (
	"fmt"
	"net/http"
	"todo/be/db"
	"todo/be/router"
	"todo/be/utils"

	"github.com/gorilla/mux"
)

func main() {
	if !db.InitDb() {
		fmt.Println("Error init database")
	}
	muxRouter := mux.NewRouter().StrictSlash(true)
	router.InitRouter(muxRouter)
	fmt.Println("http server start on :8008")
	err := http.ListenAndServe(":8008", muxRouter)
	if utils.IsError(err) {
		fmt.Println("Error when create server at :8008")
	}
}
