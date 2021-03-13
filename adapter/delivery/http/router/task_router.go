package router

import (
	"github.com/go-chi/chi"
	httpTransport "github.com/go-kit/kit/transport/http"

	"github.com/valonekowd/togo/adapter/delivery/http/decoder"
	"github.com/valonekowd/togo/adapter/delivery/http/encoder"
	"github.com/valonekowd/togo/adapter/endpoint"
	"github.com/valonekowd/togo/infrastructure/validator"
)

func NewTaskRouter(e endpoint.TaskEndpoint, options []httpTransport.ServerOption, v validator.Validator) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/", httpTransport.NewServer(
			e.Get,
			decoder.ValidatingMiddleware(decoder.GetTasks, v.Struct),
			encoder.EncodeResponse,
			options...,
		).ServeHTTP)

		r.Post("/", httpTransport.NewServer(
			e.Create,
			decoder.ValidatingMiddleware(decoder.CreateTask, v.Struct),
			encoder.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
