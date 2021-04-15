package rest

import (
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/util"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Delivery struct {
	taskService TaskService
	restUtil    util.RestUtil
}

func NewDelivery(taskService TaskService) *Delivery {
	return &Delivery{taskService: taskService, restUtil: util.NewRestUtil()}
}

//go:generate mockgen -package mock -destination mock/task_mock.go github.com/manabie-com/togo/internal/app/task/transport/rest TaskService
type TaskService interface {
	RetrieveTasks(ctx context.Context, createdDate string) ([]model.Task, error)
	AddTask(ctx context.Context, task model.Task) (model.Task, error)
}

func (d Delivery) RetrieveTasks(resp http.ResponseWriter, req *http.Request) {
	createdDate := req.FormValue("created_date")
	err := util.ValidateCreatedDate(createdDate)
	if err != nil {
		log.Error().Err(err).Msg("invalid created_date")
		d.restUtil.WriteFailedResponse(resp, http.StatusBadRequest, err)
		return
	}

	tasks, err := d.taskService.RetrieveTasks(req.Context(), createdDate)
	if err != nil {
		d.restUtil.WriteFailedResponse(resp, http.StatusInternalServerError, err)
		return
	}
	d.restUtil.WriteSuccessfulResponse(resp, tasks)
}

func (d Delivery) AddTask(resp http.ResponseWriter, req *http.Request) {
	var t model.Task
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		log.Error().Msg("invalid request body")
		d.restUtil.WriteFailedResponse(resp, http.StatusBadRequest, err)
		return
	}
	t, err = d.taskService.AddTask(req.Context(), t)
	if err != nil {
		d.restUtil.WriteFailedResponse(resp, http.StatusInternalServerError, err)
		return
	}
	d.restUtil.WriteSuccessfulResponse(resp, t)
}
