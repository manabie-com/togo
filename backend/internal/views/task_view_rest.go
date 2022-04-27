package views

import (
	"net/http"
	"fmt"
)

func TaskViewRest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	done := make(chan bool)
	go taskViewRestRouteMethod(w, req, done)
	select {
		case <-done:
			/// done
		case <-ctx.Done():
			err := ctx.Err()
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
	}
}

func taskViewRestRouteMethod(w http.ResponseWriter, req *http.Request, done chan bool) {
	switch req.Method {
		case http.MethodGet:
			taskViewRestGet(w, req, done)
		case http.MethodPost:
			taskViewRestPost(w, req, done)
		default:
			methodNotAllowedError := http.StatusMethodNotAllowed
			http.Error(w, fmt.Sprintf(req.Method), methodNotAllowedError)
			done <- true
	}
}

func taskViewRestGet(w http.ResponseWriter, req *http.Request, done chan bool) {
	done <- true
}

func taskViewRestPost(w http.ResponseWriter, req *http.Request, done chan bool) {
	done <- true
}