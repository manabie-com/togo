package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"manabie.com/togo"
	"net/http"
)

var userCrudService togo.UserCrudService

func CreateUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).WithField("r.Body", r.Body).Error("failed to read request body in func insertStopOrderHandle")
		RespondError(w, 400, err)
		return
	}
	defer r.Body.Close()

	var rq togo.UserRequest
	if err = json.Unmarshal(b, &rq); err != nil {
		RespondWithJSON(w, 400, err)
		return
	}

	if err = validateUserRequest(rq); err != nil {
		RespondError(w, 400, err)
		return
	}

	resp, err := userCrudService.CreateUser(rq)
	if err != nil {
		RespondError(w, 500, err)
		return
	}

	RespondWithJSON(w, 200, resp)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.WithError(err).WithField("r.Body", r.Body).Error("failed to read request body in func insertStopOrderHandle")
		RespondError(w, 400, err)
		return
	}
	defer r.Body.Close()

	var rq togo.UserRequest
	if err = json.Unmarshal(b, &rq); err != nil {
		log.Error(err)
		RespondError(w, 400, err)
		return
	}

	if err = validateUserRequest(rq); err != nil {
		RespondError(w, 400, err)
		return
	}

	jwt, err := userCrudService.Login(rq)
	if err != nil {
		RespondError(w, 500, err)
		return
	}

	RespondWithJSON(w, 200, jwt)
	return
}

func validateUserRequest(rq togo.UserRequest) error {
	if rq.Username == "" {
		return errors.New("username must not be empty")
	}
	return nil
}
