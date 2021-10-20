package main

import (
	"encoding/json"
	"net/http"
)

/**
 * Status Handler for application status of our ToDo Service
 * Also for checking if the project api is working
**/
func (app *ToDoService) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.Config.Env,
		Version:     version,
	}
	//convert to json the type struct and store to js variable
	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.Logger.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	// fmt.Fprint(w, "Status")
}
