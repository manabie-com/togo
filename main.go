package main

import (
	"github.com/manabie-com/togo/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//db, err := sql.Open("sqlite3", "./data.db")
	//if err != nil {
	//	log.Fatal("error opening db", err)
	//}
	//
	//http.ListenAndServe(":5050", &services.ToDoService{
	//	JWTKey: "wqGyEBBfPK9w3Lxw",
	//	Store: &sqllite.LiteDB{
	//		DB: db,
	//	},
	//})

	cmd.Serve()
}
