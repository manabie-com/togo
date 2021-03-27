package handlers

import (
	"errors"
	"fmt"
	"github.com/manabie-com/togo/domains"
	"github.com/manabie-com/togo/pkg/core"
	"github.com/manabie-com/togo/pkg/core/servehttp"
	"github.com/manabie-com/togo/pkg/utils"
	"github.com/manabie-com/togo/usecases"
	"log"
	"net/http"
	"time"
)

type GetTasksHandler struct {
	Uc   usecases.GetTasksUseCase
	Auth core.AppAuthenticator
}

func (h *GetTasksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	input, err := h.getTasksRequest(r)
	if err != nil {
		log.Printf("Error while parse request")
		utils.ResponseErrorHandler(w, err)
		return
	}

	tasks, err := h.Uc.Execute(r.Context(), input)
	if err != nil {
		if err == domains.ErrorNotFound {
			servehttp.ResponseSuccessJSON(w, []interface{}{})
			return
		}
		utils.ResponseErrorHandler(w, err)
		return
	}

	servehttp.ResponseSuccessJSON(w, tasks)
	return
}

func (h *GetTasksHandler) getTasksRequest(r *http.Request) (usecases.GetTasksInput, error) {
	input := usecases.GetTasksInput{}

	userId, err := h.Auth.ValidateToken(r)
	if err != nil {
		return input, domains.ErrorUnAuthorized
	}

	input.UserId = userId
	params := r.URL.Query()
	createdDate, ok := params["created_date"]
	if ok {
		tmp, err := time.Parse("2006-01-02", fmt.Sprintf("%v", createdDate[0]))
		if err != nil {
			return input, errors.New("created date is invalid")
		}
		input.CreatedDate = tmp
	}

	return input, nil
}
