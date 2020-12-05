package handler

import (
	"encoding/json"
	taskStorage "github.com/HoangVyDuong/togo/internal/storages/task"
	"github.com/HoangVyDuong/togo/internal/usecase/task"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func ListTasks(taskService task.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := userIDFromCtx(r.Context())
		tasks, err := taskService.GetTasks(r.Context(), id)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(map[string][]taskStorage.Task{
			"data": tasks,
		})
	}
}

func AddTask(taskService task.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t := taskStorage.Task{}
		err := json.NewDecoder(r.Body).Decode(t)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		now := time.Now()
		userID, _ := userIDFromCtx(r.Context())
		t.ID = uuid.New().String()
		t.UserID = userID
		t.CreatedDate = now.Format("2006-01-02")

		w.Header().Set("Content-Type", "application/json")

		_, err = taskService.CreateTask(r.Context(), t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(map[string]taskStorage.Task{
			"data": t,
		})
	}
}
