package main

import (
	"bytes"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/triet-truong/todo/todo/delivery/http_handlers"
	"github.com/triet-truong/todo/todo/repository"
	"github.com/triet-truong/todo/todo/usecase"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(ctx echo.Context) error {
		resp := `{"message":"success"}`
		ctx.Response().Write(bytes.NewBufferString(resp).Bytes())
		ctx.Response().Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return nil
	})

	repo := repository.NewTodoMysqlRepository("triet_truong:pw@tcp/todo?allowNativePasswords=True&parseTime=True")
	cacheStore := repository.NewTodoRedisRepository(redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	e.Debug = true
	usecase := usecase.NewTodoUseCase(repo, cacheStore)
	handler := http_handlers.NewTodoHandler(usecase)
	e.POST("/user/todo", handler.Add)

	log.Fatal(e.Start(":9090"))

}
