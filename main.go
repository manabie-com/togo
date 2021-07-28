package main

import (
	"fmt"
	"net/http"

	"github.com/manabie-com/togo/internal/database/postgresql"
	"github.com/manabie-com/togo/internal/env"
	"github.com/manabie-com/togo/internal/router"
)

func main() {
	db := postgresql.InitDatabase()
	router := router.NewRouter(db)
	http.ListenAndServe(fmt.Sprintf(":%d", env.ServerPort()), router)
}
