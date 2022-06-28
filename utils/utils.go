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

func SuccessRespond(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	res := response{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
func FailureRespond(w http.ResponseWriter, statusCode int, message string) {
	res := response{
		Status:  "Failure",
		Message: message,
		Data:    nil,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
