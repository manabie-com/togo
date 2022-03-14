package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"togo.com/config"
	"togo.com/pkg/delivery"
	"togo.com/pkg/repository"
	"togo.com/pkg/usecase"
)

func main() {
	e := echo.New()
	//init config env
	cf, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("Load config error: ", err)
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
