package services

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/storages"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey      string
	Store       storages.StorageManager
	RateLimiter storages.RateLimiter
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.getAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	}
}

func GetTodoService(jwtKey string, store storages.StorageManager, rateLimiter storages.RateLimiter) *ToDoService {
	return &ToDoService{
		JWTKey:      jwtKey,
		Store:       store,
		RateLimiter: rateLimiter,
	}
}
