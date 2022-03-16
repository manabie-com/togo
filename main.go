package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"os"
	"togo.com/config"
	"togo.com/pkg/delivery"
	"togo.com/pkg/repository"
	"togo.com/pkg/usecase"
)

func main() {
	e := echo.New()
	//init config env
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Error load config")
	}
	cf, err := config.LoadConfig(args[1])
	if err != nil {
		log.Fatal("Error valid file path", err)
	}
	db, err := sqlx.Connect("postgres", cf.PsqlInfo())
	if err != nil {
		log.Fatalln(fmt.Sprintf("Connect db error:%s", err))
	}
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	repo := repository.NewRepository(db)
	taskUc := usecase.NewTaskUseCase(repo)
	authorizeUc := usecase.NewAuthorizeUseCase(repo)
	delivery.HttpHandel(e, taskUc, authorizeUc)
	e.Logger.Fatal(e.Start(cf.ServerPort()))
}
