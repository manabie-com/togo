package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/adapter"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	JWTKey := "wqGyEBBfPK9w3Lxw"

	jwtAdapter := adapter.NewJWTAdapter(JWTKey)
	liteDB := sqllite.NewLiteDB(db)
	userUsecase := usecase.NewUserUsecase(jwtAdapter, liteDB)
	taskUsecase := usecase.NewTaskUsecase(liteDB)

	http.ListenAndServe("localhost:5050", &services.ToDoService{
		UserUsecase: userUsecase,
		TaskUsecase: taskUsecase,
	})
}
