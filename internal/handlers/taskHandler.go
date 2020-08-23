package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	taskRepository "github.com/manabie-com/togo/internal/repository"

	entity "github.com/manabie-com/togo/internal/entities"
)

// TaskHandler handles all action with Task entity
type TaskHandler struct {
	Repo *taskRepository.TaskRepository
}

// AddTask add new task
func (handler *TaskHandler) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &entity.Task{}

	now := time.Now()
	//	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	//	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	res, err := handler.Repo.Add(t)

	//err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entity.Task{
		"data": res,
	})
}

// GetAll retrive tasks by query
func (handler *TaskHandler) GetAll(resp http.ResponseWriter, req *http.Request) {
	var createdDate = req.FormValue("created_date")

	tasks, err := handler.Repo.GetAll(createdDate)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]entity.Task{
		"data": tasks,
	})
}
