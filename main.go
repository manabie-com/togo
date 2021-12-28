package main

import (
	"database/sql"
	"log"
	"togo/api"
	db "togo/db/sqlc"
	"togo/util"

	_ "github.com/lib/pq"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(`
	  __
_/  |_  ____   ____   ____
\   __\/  _ \ / ___\ /  _ \
 |  | (  <_> ) /_/  >  <_> )
 |__|  \____/\___  / \____/
            /_____/
	`)
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
