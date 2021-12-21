package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/api/route"
	"github.com/manabie-com/togo/internal/repo"
	todorepo "github.com/manabie-com/togo/internal/repo/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	mode := os.Getenv("GIN_MODE")

	gin.SetMode(mode)
	r := gin.Default()
	c := initStore()

	log.Printf("Server Start on port: %s", port)

	// conn := initStore()
	v1 := r.Group("/v1")
	route.RegisterTodo(v1, c)
	route.RegisterSwagger(r)

	start(port, r)

}

func initStore() repo.Conn {
	//Can be a factory based on config
	s := todorepo.GetTodoSchema()
	r := &repo.InMemStorage{Schema: s}

	conn, err := r.Connect()
	if err != nil {
		panic(err)
	}
	return conn
}

func start(port string, engine *gin.Engine) {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
