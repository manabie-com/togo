package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	mode := os.Getenv("GIN_MODE")

	gin.SetMode(mode)
	r := gin.Default()

	addString := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server Start on port: %s", port)

	r.Run(addString)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}
