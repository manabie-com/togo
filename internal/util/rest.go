package util

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type RestUtil interface {
	WriteFailedResponse(resp http.ResponseWriter, httpStatus int, err error)
	WriteSuccessfulResponse(resp http.ResponseWriter, data interface{})
}

type restUtil struct{}

func NewRestUtil() RestUtil {
	return &restUtil{}
}

func (r restUtil) WriteFailedResponse(resp http.ResponseWriter, httpStatus int, err error) {
	if err == nil {
		return
	}
	r.writeResponse(resp, httpStatus, map[string]string{"error": err.Error()})
}

func (r restUtil) WriteSuccessfulResponse(resp http.ResponseWriter, data interface{}) {
	r.writeResponse(resp, http.StatusOK, map[string]interface{}{"data": data})
}

func (r restUtil) writeResponse(resp http.ResponseWriter, httpStatus int, body interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(httpStatus)
	err := json.NewEncoder(resp).Encode(body)
	if err != nil {
		log.Error().Err(err).Msg("unable to write the response")
	}
}
