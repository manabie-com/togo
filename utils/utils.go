package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string
	Message string
	Data    interface{}
}

func Respond(w http.ResponseWriter, statusCode int, status, message string, data interface{}) {
	res := response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
