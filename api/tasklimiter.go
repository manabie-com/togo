package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"manabie.com/togo"
	"net/http"
	"strconv"
)

var taskLimiterService togo.TaskLimiterService

func CreateTask(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).WithField("r.Body", r.Body).Error("failed to read request body in func insertStopOrderHandle")
		RespondError(w, 400, err)
		return
	}
	defer r.Body.Close()

	var rq togo.TaskCreationRequest
	if err = json.Unmarshal(b, &rq); err != nil {
		RespondError(w, 400, err)
		return
	}

	if err = validateTaskCreationRequest(rq); err != nil {
		RespondError(w, 400, err)
		return
	}

	res, err := taskLimiterService.CreateTask(rq)
	if err != nil {
		RespondError(w, 500, err)
		return
	}

	RespondWithJSON(w, 200, res)
	return
}

func SetTaskLimiter(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).WithField("r.Body", r.Body).Error("failed to read request body in func insertStopOrderHandle")
		RespondError(w, 400, err)
		return
	}
	defer r.Body.Close()

	var rq togo.TaskLimiterSettingRequest
	if err = json.Unmarshal(b, &rq); err != nil {
		log.Error(err)
		RespondError(w, 400, err)
		return
	}

	if err = validateTaskLimiterSetting(rq); err != nil {
		log.Error(err)
		RespondError(w, 400, err)
		return
	}

	res, err := taskLimiterService.SetTaskLimit(rq.UserId, rq.TaskId, rq.Limit)
	if err != nil {
		log.Error(err)
		RespondError(w, 500, err)
		return
	}

	RespondWithJSON(w, 200, res)
	return
}

func ReloadCache(w http.ResponseWriter, r *http.Request) {
	res, err := taskLimiterService.ReloadCache()
	if err != nil {
		log.Error(err)
		RespondError(w, 500, err)
		return
	}

	RespondWithJSON(w, 200, res)
	return
}

func DoTask(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.Header.Get(USER_HEADER))
	if err != nil {
		RespondError(w, 400, errors.Wrapf(err, "Wrong userId"))
		return
	}

	vars := mux.Vars(r)
	taskIdStr, ok := vars["taskId"]
	if !ok {
		RespondError(w, 400, errors.New("must define task name for user"))
		return
	}

	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		RespondError(w, 400, errors.New("taskId is number"))
		return
	}

	res, err := taskLimiterService.DoTask(userId, taskId)
	if err != nil {
		RespondError(w, 500, err)
		return
	}

	RespondWithJSON(w, 200, res)
	return
}

func validateTaskCreationRequest(rq togo.TaskCreationRequest) error {
	if rq.TaskName == "" {
		return errors.New(fmt.Sprintf("you mus define task name"))
	}

	return nil
}

func validateTaskLimiterSetting(rq togo.TaskLimiterSettingRequest) error {
	if rq.TaskId <= 0 {
		return errors.New(fmt.Sprintf("wrong task id %d", rq.TaskId))
	}

	if rq.UserId <= 0 {
		return errors.New(fmt.Sprintf("wrong user id %d", rq.UserId))
	}

	if rq.Limit <= 0 {
		return errors.New(fmt.Sprintf("the limit of task must greater than 0"))
	}

	return nil
}
