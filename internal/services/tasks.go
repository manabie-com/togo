package services

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/manabie-com/togo/internal/utils"
)

var dailyTaskLimit int

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *sqllite.LiteDB
}

func init() {
	dailyTaskLimit = 1
}

func (s *ToDoService) listTasksByUser(resp http.ResponseWriter, req *http.Request) {
	id, _ := UserIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		utils.Value(req, "created_date"))
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	id, _ := UserIDFromCtx(req.Context())
	now := time.Now().Format("2006-01-02")
	tasks, err1 := s.Store.RetrieveTasks1(req.Context(), id, now)

	if len(tasks) >= dailyTaskLimit {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Daily task limit exceeded!",
		})
		return
	}

	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err1.Error(),
		})
		return
	}

	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, _ := UserIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
}

func (s *ToDoService) listAllTasks(resp http.ResponseWriter, req *http.Request) {
	tasks, err := s.Store.GetListTasks(req.Context())

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}

func (s *ToDoService) updateTask(resp http.ResponseWriter, req *http.Request) {
	taskId := req.FormValue("task_id")
	t := &storages.Task{}
	t.ID = taskId
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	now := time.Now().Format("2006-01-02")

	userID, _ := UserIDFromCtx(req.Context())
	t.UpdatedBy = userID
	t.UpdatedDate = now

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.UpdateTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
	resp.Header().Set("Content-Type", "application/json")
}

func (s *ToDoService) deleteTask(resp http.ResponseWriter, req *http.Request) {
	taskId := req.FormValue("task_id")
	err := s.Store.DeleteTask(req.Context(), taskId)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"message": "Task deleted successfully",
	})
}

func (s *ToDoService) getTaskById(resp http.ResponseWriter, req *http.Request) {
	taskId := req.FormValue("task_id")
	tasks, err := s.Store.GetTaskById(req.Context(), taskId)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}
