package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/triet-truong/todo/config"
	"github.com/triet-truong/todo/todo/controller/http_handlers"
	"github.com/triet-truong/todo/todo/repository"
	"github.com/triet-truong/todo/todo/usecase"
)

func init() {
	// Load env vars
	config.Load()
}

func main() {
	e := echo.New()
	e.GET("/hello", func(ctx echo.Context) error {
		resp := `{"message":"success"}`
		ctx.Response().Write(bytes.NewBufferString(resp).Bytes())
		ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return nil
	})

	e.Debug = true
	repo := repository.NewTodoMysqlRepository(config.DBConnectionURL())
	cacheStore := repository.NewTodoRedisRepository(redis.Options{
		Addr:     config.CacheConnectioURL(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	usecase := usecase.NewTodoUseCase(repo, cacheStore)
	handler := http_handlers.NewTodoHandler(usecase)
	e.POST("/user/todo", handler.Add)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port())))

}
