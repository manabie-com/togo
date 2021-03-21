package main

import (
	"fmt"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/db"
	"github.com/manabie-com/togo/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	config.LoadEnv("")
	env := config.NewEnv

	conn := db.ConnectDB()

	defer db.DisconnectDB(conn)

	srv := http.Server{
		Addr:           "0.0.0.0:" + env.ServerPort,
		Handler:        ApplicationRecovery(routes.NewRouter(conn)),
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server is listening on port: " + env.ServerPort)
	log.Fatal(srv.ListenAndServe())
	log.Println("Shutting down...")
}

func ApplicationRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {
				fmt.Fprintln(os.Stderr, "Recovered from application error occurred")
				_, _ = fmt.Fprintln(os.Stderr, err)
				w.WriteHeader(http.StatusInternalServerError)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
