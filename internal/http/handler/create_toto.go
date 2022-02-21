package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/chi07/todo/internal/http/response"
	"github.com/chi07/todo/internal/model"
)

type CreateTaskHandler struct {
	service CreateTaskService
}

func NewCreateTaskHandler(service CreateTaskService) *CreateTaskHandler {
	return &CreateTaskHandler{service: service}
}

type CreateTaskRequest struct {
	Title string `validate:"required"`
}

var validate *validator.Validate

func (ctr CreateTaskRequest) Bind(req *http.Request) error {
	validate = validator.New()
	if err := validate.Struct(&ctr); err != nil {
		return err
	}

	return nil
}

func (h *CreateTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authID := r.Header.Get("auth-user-id")
	userID, err := uuid.Parse(authID)
	if err != nil {
		response.Error(w, r, http.StatusForbidden, http.StatusForbidden, "invalid user")
		return
	}

	req := CreateTaskRequest{}
	err = render.Bind(r, &req)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, http.StatusBadRequest, err.Error())
		return
	}
	task := &model.Task{
		UserID: userID,
		Title:  req.Title,
	}
	taskID, err := h.service.CreateTask(r.Context(), task)
	if err != nil {
		if errors.Is(err, model.ErrorNotFound) {
			response.Error(w, r, http.StatusBadRequest, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, model.ErrorNotAllowed) {
			response.Error(w, r, http.StatusForbidden, http.StatusForbidden, err.Error())
			return
		}

		response.Error(w, r, http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]interface{}{"taskID": taskID}

	response.Success(w, r, http.StatusCreated, data)
}
