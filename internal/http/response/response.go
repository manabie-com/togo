package response

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type Response struct {
	Error *ErrorResponse `json:"error,omitempty"`
	Data  interface{}    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (r *Response) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

func Error(w http.ResponseWriter, r *http.Request, httpCode int, errCode int, message string) {
	render.Status(r, httpCode)
	err := render.Render(w, r, &Response{
		Error: &ErrorResponse{
			Code:    errCode,
			Message: message,
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Cannot render error response.")
	}
	return
}

func Success(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	render.Status(r, code)
	err := render.Render(w, r, &Response{
		Data: data,
	})
	if err != nil {
		log.Error().Err(err).Interface("data", data).Msg("Cannot render success response.")
	}
}
