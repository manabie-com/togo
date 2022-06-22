package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	status  string
	message string
	data    interface{}
}

func Respond(w http.ResponseWriter, statusCode int, status, message string, data interface{}) {
	res := response{
		status:  status,
		message: message,
		data:    data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
