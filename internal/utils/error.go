package utils

import (
	"encoding/json"
	"net/http"
)

type NotFoundError struct {
}

func (e *NotFoundError) Error() string {
	return "NotFoundID"
}

func HttpResponseInternalServerError(response http.ResponseWriter, error string) {
	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(map[string]string{
		"error": error,
	})
}

func HttpResponseBadRequest(response http.ResponseWriter, error string) {
	response.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(response).Encode(map[string]string{
		"error": error,
	})
}

func HttpResponseUnauthorized(response http.ResponseWriter, error string) {
	response.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(response).Encode(map[string]string{
		"error": error,
	})
}
