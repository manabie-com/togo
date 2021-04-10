package main

import (
	"context"
	"log"
	"os"

	"github.com/manabie-com/togo/internal/infra/app"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//db, err := sql.Open("sqlite3", "./data.db")
	//if err != nil {
	//	log.Fatal("error opening db", err)
	//}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	ctx := context.Background()
	cli, cleanup, err := app.InitApplication(ctx)
	if err != nil {
		panic(err)
	}
	app.HandleSigterm(cleanup)

	err = cli.Commands().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
