package services

import (
	"encoding/json"
	"net/http"
)

func (s *ToDoService) ping(resp http.ResponseWriter, req *http.Request) {
	json.NewEncoder(resp).Encode(map[string]string{
		"data": "Pong",
	})
}

