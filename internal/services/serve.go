package services

import (
	"log"
	"net/http"
)

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}
	path := req.URL.Path
	method := req.Method
	if path != "/login" && path != "/register" {
		var ok bool
		req, ok = s.ValidToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	switch path {
	case "/login":
		if method == http.MethodPost {
			s.GetAuthToken(resp, req)
			return
		}
	case "/register":
		if method == http.MethodPost {
			s.RegisterUser(resp, req)
		}
	case "/tasks":
		switch method {
		case http.MethodGet:
			s.listAllTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	case "/tasks/manage":
		switch method {
		case http.MethodGet:
			s.listTasksByUser(resp, req)
		case http.MethodPut:
			s.updateTask(resp, req)
		case http.MethodDelete:
			s.deleteTask(resp, req)
		}
	case "/task12":
		if method == http.MethodGet {
			s.getTaskById(resp, req)
		}
	case "/users":
		if method == http.MethodGet {
			s.getUsers(resp, req)
		}
	}

}
