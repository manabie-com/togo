package middlewares

import (
	"github.com/go-chi/cors"
	"net/http"
	"togo/config"
)

func (m middleware) WithCors() func(http.Handler) http.Handler {
	conf := config.GetConfig()
	return cors.New(cors.Options{
		AllowedOrigins:   conf.Cors.AllowedOrigins,
		AllowedMethods:   conf.Cors.AllowedMethods,
		AllowedHeaders:   conf.Cors.AllowedHeaders,
		AllowCredentials: true,
	}).Handler
}
