package services

import (
	"log"
	"net/http"
)

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	path := req.URL.Path
	method := req.Method
	if method == http.MethodOptions {
		response(&resp, http.StatusOK, nil)
		return
	}

	if path == "/login" && method == http.MethodGet {
		ctx := req.Context()
		id := value(req, "user_id")
		pwd := value(req, "password")
		token, err := s.getAuthToken(ctx, id, pwd)

		body := make(map[string]string)
		if err != nil {
			body["error"] = err.Error()
			response(&resp, http.StatusUnauthorized, body)
		} else {
			body["token"] = token
			response(&resp, http.StatusOK, body)
		}
		return
	}

	var err error
	req, err = s.parseUserIdToContext(req)
	if err != nil {
		response(&resp, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	switch path {
	case "/tasks":
		switch method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		default:
			response(&resp, http.StatusNotFound, nil)
		}
		return
	default:
		response(&resp, http.StatusNotFound, nil)
	}

}
