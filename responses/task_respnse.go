package responses

import (
	"encoding/json"
	"net/http"
)

type TaskResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := UserResponse{Status: status, Message: "success", Data: res}
	json.NewEncoder(w).Encode(response)
}
