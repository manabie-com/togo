package main

import (
	"fmt"
	"net/http"
	"os"

	lr "togo/utils/logger"

	router "togo/http"
)

var (
	httpRouter router.Router = router.NewChiRouter()
)

func main() {
	// Set logging
	logger := lr.NewLogger(os.Getenv("LOG_LEVEL"))
	port := os.Getenv("PORT")

	httpRouter.POST("/tasks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	logger.Info().Msgf("Serving at %v", port)
	httpRouter.SERVE(port)
}
