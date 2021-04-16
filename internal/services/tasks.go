package services

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/manabie-com/togo/internal/storages/entities"
)

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*entities.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &entities.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	resp.Header().Set("Content-Type", "application/json")

	isAllowed, err := s.RateLimiter.Allow(req.Context(), userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"data": "an internal server error occurred",
		})
		return
	}

	if !isAllowed {
		resp.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(resp).Encode(map[string]string{
			"data": "user has created max task for today",
		})
		return
	}

	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entities.Task{
		"data": t,
	})
}
