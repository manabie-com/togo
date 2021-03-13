package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kitTransport "github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/rs/cors"

	"github.com/valonekowd/togo/adapter/delivery/http/decoder"
	"github.com/valonekowd/togo/adapter/delivery/http/encoder"
	"github.com/valonekowd/togo/adapter/delivery/http/option"
	"github.com/valonekowd/togo/adapter/delivery/http/router"
	"github.com/valonekowd/togo/adapter/endpoint"
	"github.com/valonekowd/togo/infrastructure/validator"
)

const (
	apiVersion0 = ""
)

func NewHTTPHandler(e endpoint.ServerEndpoint, v validator.Validator, logger log.Logger) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	var r chi.Router
	{
		r = chi.NewRouter()
		r.Use(c.Handler)
	}

	options := []httpTransport.ServerOption{
		httpTransport.ServerBefore(
			option.LogRequestInfo(logger),
			jwt.HTTPToContext(),
		),
		httpTransport.ServerErrorHandler(kitTransport.NewLogErrorHandler(logger)),
		httpTransport.ServerErrorEncoder(encoder.EncodeError),
	}

	r.Get("/", httpTransport.NewServer(
		e.HealthCheck,
		httpTransport.NopRequestDecoder,
		httpTransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/api", func(r chi.Router) {
		r.Route("/"+apiVersion0, func(r chi.Router) {
			r.Post("/sign-in", httpTransport.NewServer(
				e.User.SignIn,
				decoder.ValidatingMiddleware(decoder.SignIn, v.Struct),
				encoder.EncodeResponse,
				options...,
			).ServeHTTP)

			r.Post("/sign-up", httpTransport.NewServer(
				e.User.SignUp,
				decoder.ValidatingMiddleware(decoder.SignUp, v.Struct),
				encoder.EncodeResponse,
				options...,
			).ServeHTTP)

			r.Route("/tasks", router.NewTaskRouter(e.Task, options, v))
		})
	})

	return r
}
