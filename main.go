package main

import (
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/delivery/rest"
	"github.com/manabie-com/togo/internal/pkgs/clients"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/routers"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/tokens"
	"go.uber.org/zap"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

}

func main() {
	conf, err := config.Load()
	if err != nil {
		zap.L().Panic("load config error", zap.Error(err))
	}
	//dbConf := clients.PSQLConfig{
	//	DSN: "host=localhost user=togo password=ad34a$dg dbname=manabie_togo port=5432 sslmode=disable",
	//}
	db, err := clients.InitPSQLDB(conf.DB)
	if err != nil {
		log.Fatal("error opening db", err)
	}



	taskService := repositories.NewTaskRepo(db)
	userService := repositories.NewUserRepo(db)
	tokenService := tokens.NewTokenManager("wqGyEBBfPK9w3Lxw", userService)

	todoService := services.NewToDoService(taskService, tokenService)
	serializer := rest.NewSerializer(todoService)

	server := routers.NewServer(serializer, conf.HTTP)

	server.Run()
}
