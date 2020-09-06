package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/manabie-com/togo/internal/repository"

	entity "github.com/manabie-com/togo/internal/entities"

	"github.com/manabie-com/togo/internal/services"
)

// TaskHandler handles all action with Task entity
type TaskHandler struct {
	//Repo *taskRepository.TaskRepository
}

// AddTask add new task
func (handler *TaskHandler) AddTask(resp http.ResponseWriter, req *http.Request) {
	task := &entity.Task{}

	var res map[string]interface{}

	json.NewDecoder(req.Body).Decode(&res)

	userID, _ := services.UserIDFromCtx(req.Context())

	tasksOfCurrentUser, _ := repository.TaskRepo.GetByUserID(userID, time.Now().Format("2006-01-02"))

	if len(tasksOfCurrentUser) > 5 {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Reach out the daily limit",
		})
		return
	}

	task.UserID = userID

	task.Content = res["content"].(string)

	result, err := repository.TaskRepo.Add(task)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entity.Task{
		"data": result,
	})
}

// GetAll retrive tasks by query
func (handler *TaskHandler) GetAll(resp http.ResponseWriter, req *http.Request) {
	var createdDate = req.FormValue("created_date")

	tasks, err := repository.TaskRepo.GetAll(createdDate)

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

// GetByID get task by id
func (handler *TaskHandler) GetByID(resp http.ResponseWriter, req *http.Request) {
	var id = mux.Vars(req)["id"]

	log.Printf("Get task by id %s \n", id)

	task, err := repository.TaskRepo.GetByID(id)

	if task.ID == "" {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Entity not Found",
		})
		return
	}

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entity.Task{
		"data": task,
	})
}
