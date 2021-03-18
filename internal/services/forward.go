package services

import (
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/storages"
	"log"
	"net/http"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *storages.LiteDB
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		writeHeader(&resp, http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.GetAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			writeHeader(&resp, http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.ListTasks(resp, req)
		case http.MethodPost:
			s.AddTask(resp, req)
		}
		return
	}
}

func Value(v string) sql.NullString {
	return sql.NullString{
		String: v,
		Valid:  true,
	}
}

func responseJson(resp *http.ResponseWriter, mJson interface{}) {
	json.NewEncoder(*resp).Encode(mJson)
}

func writeHeader(resp *http.ResponseWriter, statusCode int) {
	(*resp).WriteHeader(statusCode)
}
