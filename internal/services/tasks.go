package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/entities"
)

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	date, err := time.Parse("2006-01-02", req.URL.Query().Get("created_date"))
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "date should follow the format YYYY-MM-DD",
		})
		return
	}

	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		id,
		date,
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*entities.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var task entities.Task
	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "invalid request body",
		})
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	resp.Header().Set("Content-Type", "application/json")

	isAllowed, err := s.RateLimiter.Allow(req.Context(), userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "an internal server error occurred",
		})
		return
	}

	if !isAllowed {
		resp.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "user has created max task for today",
		})
		return
	}

	task = entities.Task{
		ID:          uuid.New().String(),
		UserID:      userID,
		Content:     task.Content,
		CreatedDate: now.Format("2006-01-02"),
	}
	err = s.Store.AddTask(req.Context(), &task)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"message": "failed to add task",
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entities.Task{
		"data": &task,
	})
}
