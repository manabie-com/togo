package main

import (
	"net/http"

	"github.com/manabie-com/togo/internal/routers"
)

func main() {

	// create interface for db

	// db, err := sql.Open("sqlite3", "./data.db")
	// if err != nil {
	// 	log.Fatal("error opening db", err)
	// }


	// http.ListenAndServe(":5050", &services.ToDoService{
	// 	JWTKey: "wqGyEBBfPK9w3Lxw",
	// 	Store: &sqllite.LiteDB{
	// 		DB: db,
	// 	},
	// })

	http.ListenAndServe(":5050", routers.OhRouter().InitRouter())

}
