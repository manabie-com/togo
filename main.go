package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	mode := os.Getenv("GIN_MODE")

	gin.SetMode(mode)
	r := gin.Default()

	log.Printf("Server Start on port: %s", port)

	v1 := r.Group("/v1")
	route.RegisterTodo(v1)

	start(port, r)

}

func start(port string, engine *gin.Engine) {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: engine,
	}

	srv.ListenAndServe()
}
