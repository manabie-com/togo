package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
