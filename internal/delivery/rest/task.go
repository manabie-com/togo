package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

type Serializer struct {
	TodoService services.TodoService
}

func NewSerializer(todoService services.TodoService) *Serializer {
	return &Serializer{
		TodoService: todoService,
	}
}

func (s *Serializer) ListTasks(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	userID, _ := utils.UserIDFromCtx(req.Context())
	createdAt := utils.Value(req, "created_date")
	tasks, err := s.TodoService.ListTasks(userID, createdAt)
	if err != nil {
		zap.L().Error("list tasks error.", zap.Error(err))
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
		zap.L().Error("add tasks error.", zap.Error(err))
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

	task, err := s.TodoService.AddTask(userID, &t)
	if err != nil {
		zap.L().Error("add tasks error.", zap.Error(err))
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
