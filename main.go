package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/internal/routers"
	"github.com/manabie-com/togo/pkg/postgresql"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

// enviroment value
var (
	port string
	dsn  string
)

func init() {
	godotenv.Load()
	port = os.Getenv("PORT")
	dsn = os.Getenv("DSN")
}

func main() {
	db, err := postgresql.GetInstance(dsn)
	if err != nil {
		log.Fatal("error to connecting db", err)
	}

	routersInit := routers.InitRouter(db)
	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	logrus.Infof("server start http server listening %v", port)

	server.ListenAndServe()

}
