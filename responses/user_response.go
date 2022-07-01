package responses

import (
	"encoding/json"
	"net/http"
)

type UserResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}

func WriteResponseUser(w http.ResponseWriter, token string, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := UserResponse{Status: status, Message: "success", Token: token, Data: res}
	json.NewEncoder(w).Encode(response)
}
