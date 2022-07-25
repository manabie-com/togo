package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/handler"
	"github.com/huuthuan-nguyen/manabie/app/middleware"
	"github.com/huuthuan-nguyen/manabie/config"
)

func NewRouter(config *config.Config, handler *handler.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Routes for api
	router = SetAPIRoutes(router, config, handler)

	recoveryMiddleware := middleware.RecoveringMiddleware()
	loggingMiddleware := middleware.LoggingMiddleware()
	compressMiddleware := handlers.CompressHandler
	rateLimitingMiddleware := middleware.RateLimitingMiddleware()
	corsMiddleware := middleware.CORSMiddleware()
	prometheusMiddleware := middleware.PrometheusMiddleware()
	validateMiddleware := middleware.ValidateMiddleware()
	router.Use(
		compressMiddleware,
		loggingMiddleware,
		recoveryMiddleware,
		prometheusMiddleware,
		corsMiddleware,
		rateLimitingMiddleware,
		validateMiddleware,
	)

	return router
}
