package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manabie/project/config"
	todoHttp "github.com/manabie/project/internal/http"
	todoRepository "github.com/manabie/project/internal/repository"
	todoRouter "github.com/manabie/project/internal/router"
	todoServer "github.com/manabie/project/internal/server"
	todoUsecase "github.com/manabie/project/internal/usecase"
	"github.com/manabie/project/middleware"
	"github.com/manabie/project/model"
	"github.com/manabie/project/pkg/jwt"
	"github.com/manabie/project/pkg/postgres"
	"github.com/manabie/project/pkg/snowflake"
	"github.com/manabie/project/pkg/hash"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	accessControl := middleware.NewAccessController(jwt)
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)
	todoServer := todoServer.NewServer(todoRouter, accessControl, router)
	if err := todoServer.RunServer(); err != nil {
		log.Println("run the api of task todo user ", err)
	}

	if err := router.Run(conf.Server.PortServer); err != nil {
		log.Println("run port server failed port %s: ",conf.Server.PortServer)
	}
}