package main

import (
	"fmt"
	"github.com/manabie-com/togo/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	srv := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: ApplicationRecovery(routes.Router),
		//ReadTimeout:  15 * time.Second,
		//WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
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
