package transport

import (
	"encoding/json"
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/api/task/storages"
	"github.com/manabie-com/togo/internal/api/task/usecase"
	"github.com/manabie-com/togo/internal/api/utils"
	"github.com/manabie-com/togo/internal/pkg/token"
	"net/http"
)

const (
	CreatedDateField = "created_date"
	PageField        = "page"
	Limit            = "limit"
)

type Task struct {
	TaskUC usecase.Task
}

func (t *Task) List(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	userID, err := utils.GetValueFromCtx(ctx, token.UserIDField)
	if err != nil {
		utils.WriteJSON(ctx, resp, http.StatusInternalServerError, nil, err)
		return
	}

	createdDate := req.FormValue(CreatedDateField)
	if createdDate == "" {
		utils.WriteJSON(ctx, resp, http.StatusBadRequest, nil, errors.New(dictionary.CreatedDateRequestEmpty))
		return
	}

	if ok := utils.ValidateDateFromString(createdDate, utils.DefaultLayout); !ok {
		utils.WriteJSON(ctx, resp, http.StatusBadRequest, nil, errors.New(dictionary.DateRequestEmptyIsNotValid))
		return
	}

	page, limit := 0, 0
	pageStr := req.FormValue(PageField)
	if value, ok := utils.ValidateInputIsInteger(pageStr); pageStr != "" && ok {
		page = value
	}

	limitStr := req.FormValue(Limit)
	if value, ok := utils.ValidateInputIsInteger(limitStr); limitStr != "" && ok {
		limit = value
	}

	tasks, err := t.TaskUC.List(ctx, userID, createdDate, page, limit)
	if err != nil {
		utils.WriteJSON(ctx, resp, http.StatusInternalServerError, nil, err)
		return
	}

	utils.WriteJSON(ctx, resp, http.StatusOK, tasks, err)
}

func (t *Task) Add(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	task := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(task)
	if err != nil {
		utils.WriteJSON(ctx, resp, http.StatusBadRequest, nil, err)
		return
	}

	valErr := utils.ValidateRequest(task)
	if valErr != nil {
		utils.WriteJSON(ctx, resp, http.StatusBadRequest, nil, valErr)
		return
	}

	userID, err := utils.GetValueFromCtx(ctx, token.UserIDField)
	if err != nil {
		utils.WriteJSON(ctx, resp, http.StatusInternalServerError, nil, err)
		return
	}

	task, err = t.TaskUC.Add(ctx, userID, task)
	if err != nil {
		if err.Error() == errors.New(dictionary.UserReachTaskLimit).Error() {
			utils.WriteJSON(ctx, resp, http.StatusForbidden, nil, err)
			return
		}
		utils.WriteJSON(ctx, resp, http.StatusInternalServerError, nil, err)
		return
	}

	utils.WriteJSON(ctx, resp, http.StatusOK, task, nil)
}
