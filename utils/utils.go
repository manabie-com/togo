package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, status int, payload map[string]interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
