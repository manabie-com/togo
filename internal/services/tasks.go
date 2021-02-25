package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
)

func (s *ToDoService) setHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")
		resp.Header().Set("Access-Control-Allow-Methods", "*")

		if req.Method == http.MethodOptions {
			resp.WriteHeader(http.StatusOK)
			return
		}

		next(resp, req)
	}
}

func (s *ToDoService) tasksHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Println(req.Method, req.URL.Path)
		switch req.Method {
		case http.MethodPost:
			s.addTaskHandler(resp, req)
		case http.MethodGet:
			s.listTasksHandler(resp, req)
		default:
			resp.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (s *ToDoService) listTasksHandler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	id, _ := userIDFromCtx(req.Context())

	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
	)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		respData := map[string]string{
			"error": err.Error(),
		}
		if err = json.NewEncoder(resp).Encode(respData); err != nil {
			log.Println(err)
		}
		return
	}

	respData := map[string][]*storages.Task{
		"data": tasks,
	}
	if err = json.NewEncoder(resp).Encode(respData); err != nil {
		log.Println(err)
	}
}

func (s *ToDoService) addTaskHandler(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, ok := userIDFromCtx(req.Context())
	if !ok {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = time.Now().Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		respData := map[string]string{
			"error": err.Error(),
		}
		if err := json.NewEncoder(resp).Encode(respData); err != nil {
			log.Println(err)
		}
		return
	}

	respData := map[string]*storages.Task{
		"data": t,
	}
	if err := json.NewEncoder(resp).Encode(respData); err != nil {
		log.Println(err)
	}
}
