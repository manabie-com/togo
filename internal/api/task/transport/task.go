package transport

import (
	"errors"
	"github.com/manabie-com/togo/internal/api/dictionary"
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

}
