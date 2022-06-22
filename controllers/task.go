package controllers

import (
	"net/http"

	u "github.com/manabie-com/togo/utils"
)

var GetTasks = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}

var GetTask = func(w http.ResponseWriter, r *http.Request) {
	// check taskID exist
	// if not exists
	u.Respond(w, http.StatusNotFound, map[string]interface{}{})
	// else
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}

var Add = func(w http.ResponseWriter, r *http.Request) {
	// check task number today greater than or equal to current user limitDayTasks
	// if true
	u.Respond(w, http.StatusNotAcceptable, map[string]interface{}{})
	// else
	// update task number today += 1
	u.Respond(w, http.StatusCreated, map[string]interface{}{})
}

var Edit = func(w http.ResponseWriter, r *http.Request) {
	// check taskID exist
	// if not exists
	u.Respond(w, http.StatusNotFound, map[string]interface{}{})
	// else
	u.Respond(w, http.StatusOK, map[string]interface{}{})
}
