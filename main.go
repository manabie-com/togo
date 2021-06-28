package main

import (
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/delivery/rest"
	"github.com/manabie-com/togo/internal/pkgs/clients"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/routers"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/tokens"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
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
	db, err := clients.InitPSQLDB(conf.DB)
	if err != nil {
		zap.L().Panic("connecting to db error", zap.Error(err))
	}

	cacheClient, err := clients.InitRedisClient(conf.Redis)
	if err != nil {
		zap.L().Panic("connecting to redis error", zap.Error(err))
	}

	defer cacheClient.Close()

	taskService := repositories.NewTaskRepo(db)
	userService := repositories.NewUserRepo(db)
	tokenService := tokens.NewTokenManager(conf.Token.JWT, userService)
	cachingService := repositories.NewCacheManager(cacheClient)


	todoService := services.NewToDoService(taskService, userService, tokenService, cachingService)
	serializer := rest.NewSerializer(todoService)

	server := routers.NewServer(serializer, conf.HTTP)

	server.Run()
}
