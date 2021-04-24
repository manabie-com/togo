package routes

import (
	"net/http"

	"github.com/manabie-com/togo/internal/app/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Path         string
	Method       string
	Handler      http.HandlerFunc
	AuthRequired bool
}

func Install(r *mux.Router, togoRoutes TogoRoutes) {
	for _, route := range togoRoutes.Routes() {
		if route.AuthRequired {
			r.HandleFunc(route.Path,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.CORSMiddleware(
							middlewares.SetMiddlewareAuthentication(route.Handler),
						),
					),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Path,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.CORSMiddleware(route.Handler),
					),
				),
			).Methods(route.Method)
		}
	}
}
