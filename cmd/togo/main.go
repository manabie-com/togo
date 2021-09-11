package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/time/rate"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/sqlite"
)

var (
	requestsPerDay = flag.Int64("requests-per-day", 5, "Number of requests per day")
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}

	liteDB := sqlite.NewLiteDB(db)
	userLimiter := services.NewUserRateLimiter(rate.Every(24 * time.Hour / time.Duration(*requestsPerDay)), int(*requestsPerDay))
	todoService := services.NewToDoService("wqGyEBBfPK9w3Lxw", liteDB, userLimiter)

	http.ListenAndServe(":5050", todoService)
}
