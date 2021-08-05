package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/delivery/http"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	db2 "github.com/manabie-com/togo/pkg/db"
	"github.com/manabie-com/togo/pkg/model"
)

func main() {
	app := initHttp()
	app.Run(":5050")
}

func initHttp() *gin.Engine {

	// connect to database
	connStr := `host=postgres port=5432 user=admin password=admin dbname=todo sslmode=disable binary_parameters=yes`
	db, err := db2.NewDatabase("postgres", connStr, 10, 2)
	if err != nil {
		panic("no db connection")
	}

	// migration data
	db2.Migration(db, []interface{}{&model.Task{}, &model.User{}})

	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(gin.Recovery())
	r.Use(cors.New(config))

	repository, _ := postgres.NewRepository(db)
	userService := services.NewUserService(repository, "wqGyEBBfPK9w3Lxw")
	taskSerVice := services.NewTaskService(repository)

	// create first user , dont care about error
	userService.CreateUser(&model.User{
		UserName: "firstUser",
		Password: "example",
		MaxTodo:  5,
	})

	handler := http.NewHandler(
		http.WithTaskService(taskSerVice),
		http.WithUserService(userService),
		http.WithGinEngine(r),
	)
	handler.InitRoute()

	return r
}
