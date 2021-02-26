package services

import (
	"encoding/json"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/storages"
)

func (s *ToDoService) setHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")
		resp.Header().Set("Access-Control-Allow-Methods", "*")

		resp.Header().Set("Content-Type", "application/json")

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
	id, _ := userIDFromCtx2(req.Context())

	createdDate, err := time.Parse("2006-01-02", req.FormValue("created_date"))
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	tasks, err := s.pg.GetTasks(req.Context(), id, createdDate)
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

	if err = json.NewEncoder(resp).Encode(newDataResp(tasks)); err != nil {
		log.Println(err)
	}
}

func (s *ToDoService) addTaskHandler(resp http.ResponseWriter, req *http.Request) {
	defer func() {
		_ = req.Body.Close()
	}()

	task := &storages.PgTask{}
	err := json.NewDecoder(io.LimitReader(req.Body, maxJsonSize)).Decode(task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, ok := userIDFromCtx2(req.Context())
	if !ok {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	task.UsrId = userID
	task.CreateAt = time.Now()

	switch err := s.pg.InsertTask(req.Context(), task); err {
	case nil:
		if err := json.NewEncoder(resp).Encode(newDataResp(task)); err != nil {
			log.Println(err)
		}
	case postgres.ErrUserMaxTodoReached:
		resp.WriteHeader(http.StatusTooManyRequests)
		if err := json.NewEncoder(resp).Encode(newErrResp(err.Error())); err != nil {
			log.Println(err)
		}
	default:
		resp.WriteHeader(http.StatusInternalServerError)
	}
}
