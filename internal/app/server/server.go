package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manabie-com/togo/internal/app/models"
)

func Run() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5050"
	}

	db := models.Connect()
	if db != nil {
		defer db.Close()
	}

	models.TestConnection()

	fmt.Printf("Api running on port %s\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
