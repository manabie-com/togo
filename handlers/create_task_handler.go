package handlers

import (
	"encoding/json"
	"errors"
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/pkg/core"
	"github.com/manabie-com/togo/pkg/core/servehttp"
	"github.com/manabie-com/togo/pkg/utils"
	"github.com/manabie-com/togo/usecases"
	"log"
	"net/http"
)

type CreateTaskHandler struct {
	Uc   usecases.CreateTaskUseCase
	Auth core.AppAuthenticator
}

func (h *CreateTaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	input, err := h.getTaskRequest(r)
	if err != nil {
		log.Printf("Error while parse request")
		utils.ResponseErrorHandler(w, err)
		return
	}

	task, err := h.Uc.Execute(r.Context(), *input)
	if err != nil {
		log.Printf("Error while create task")
		utils.ResponseErrorHandler(w, err)
		return
	}

	servehttp.ResponseSuccessJSON(w, task)
	return
}

func (h *CreateTaskHandler) getTaskRequest(r *http.Request) (*usecases.TaskInput, error) {
	input := &usecases.TaskInput{}

	userId, err := h.Auth.ValidateToken(r)
	if err != nil {
		return nil, domains.ErrorUnAuthorized
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		return nil, errors.New("cannot parse create task request")
	}

	input.UserId = userId

	validator, _ := utils.NewGoPlayground()
	err = validator.Validate(input)
	if err != nil {
		return nil, err
	}

	if len(validator.Messages()) > 0 {
		return nil, errors.New(validator.Messages()[0])
	}

	return input, nil
}
