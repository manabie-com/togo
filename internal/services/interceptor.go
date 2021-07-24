package services

import (
	"encoding/json"
	"net/http"
)

func (c httpInterceptor) WithHandler(handler httpHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		writer.Header().Set("Content-Type", "application/json")

		if req.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}

		resp, err := c(req, handler)
		if err != nil {
			var statusCode = http.StatusInternalServerError
			if serviceErr, ok := err.(*serviceError); ok {
				statusCode = serviceErr.StatusCode()
			}
			writer.WriteHeader(statusCode)

			json.NewEncoder(writer).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}
		json.NewEncoder(writer).Encode(map[string]interface{}{
			"data": resp,
		})
	}
}

func NewInterceptor(interceptors ...httpInterceptor) httpInterceptor {
	n := len(interceptors)

	return func(req *http.Request, handler httpHandler) (resp interface{}, err error) {
		chainer := func(currentInter httpInterceptor, currentHandler httpHandler) httpHandler {
			return func(currentReq *http.Request) (resp interface{}, err error) {
				return currentInter(currentReq, currentHandler)
			}
		}
		chainedHandler := handler

		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}
		return chainedHandler(req)
	}
}

type httpInterceptor func(req *http.Request, handler httpHandler) (resp interface{}, err error)

type httpHandler func(req *http.Request) (resp interface{}, err error)