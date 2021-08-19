package transports

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/models"
	repository "github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/utils"
)

// ToDoService implement HTTP server
type ToDoTaskController struct {
	service *services.ToDoTaskService
}

func NewToDoTaskController(db *repository.DB, jwtKey string) *ToDoTaskController {
	return &ToDoTaskController{
		service: &services.ToDoTaskService{
			JWTKey: jwtKey,
			Store:  db,
		},
	}
}

func (t *ToDoTaskController) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := utils.NewJwtUtil(t.service.JWTKey).UerIDFromCtx(req.Context())
	tasks, err := t.service.ListTasks(req.Context(), id, req.FormValue("created_date"))

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*models.Task{
		"data": tasks,
	})
}

func (s *ToDoTaskController) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &models.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err = s.service.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusNotAcceptable)
		return
	}

	json.NewEncoder(resp).Encode(map[string]*models.Task{
		"data": t,
	})
}
