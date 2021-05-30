package services

import (
	"encoding/json"
	"net/http"
)

func responseOK(resp http.ResponseWriter, v interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": v,
	})
}

func responseError(resp http.ResponseWriter, statusCode int, err interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	json.NewEncoder(resp).Encode(map[string]interface{}{
		"error": err,
	})
}
