package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/utils"
	"net/http"
)

type Serializer struct {
	TodoService *services.ToDoService
}

func NewSerializer(todoService *services.ToDoService) *Serializer {
	return &Serializer{
		TodoService: todoService,
	}
}

func (s *Serializer) ListTasks(resp http.ResponseWriter, req *http.Request) {
	userID, _ := utils.UserIDFromCtx(req.Context())
	createdAt := utils.Value(req, "created_date")
	tasks, err := s.TodoService.ListTasks(userID, createdAt)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(
			utils.BuildErrorResponseRequest(&utils.Meta{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		return
	}

	json.NewEncoder(resp).Encode(
		utils.BuildSuccessResponseRequest(&utils.Meta{
			Code:    http.StatusOK,
			Message: utils.SuccessRequestMessage,
		}, tasks))
}

func (s *Serializer) AddTask(resp http.ResponseWriter, req *http.Request) {
	var t repositories.Task
	err := json.NewDecoder(req.Body).Decode(&t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(
			utils.BuildErrorResponseRequest(&utils.Meta{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		return
	}

	userID, _ := utils.UserIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID

	resp.Header().Set("Content-Type", "application/json")

	task, err := s.TodoService.AddTask(&t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(
			utils.BuildErrorResponseRequest(&utils.Meta{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		return
	}

	json.NewEncoder(resp).Encode(
		utils.BuildSuccessResponseRequest(&utils.Meta{
			Code:    http.StatusCreated,
			Message: utils.SuccessRequestMessage,
		}, task))
}
