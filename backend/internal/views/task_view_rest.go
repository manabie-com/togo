package views

import (
	"net/http"
	"fmt"
	"manabie.com/internal/controllers"
    "encoding/json"
	"context"
	"strconv"
)

type TaskViewRest struct {
	controller *controllers.TaskController
}

func MakeTaskViewRest(
	iController *controllers.TaskController,
) *TaskViewRest {
	return &TaskViewRest {
		controller: iController,
	}
}

func (r *TaskViewRest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("serving")
	ctx := req.Context()
	done := make(chan bool)
	go r.taskViewRestRouteMethod(w, req, done)
	select {
		case <-done:
			/// done
		case <-ctx.Done():
			err := ctx.Err()
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
	}
}

func (r *TaskViewRest) taskViewRestRouteMethod(w http.ResponseWriter, req *http.Request, done chan bool) {
	switch req.Method {
		case http.MethodGet:
			r.taskViewRestGet(w, req, done)
		case http.MethodPost:
			r.taskViewRestPost(w, req, done)
		default:
			http.Error(w, fmt.Sprintf(req.Method), http.StatusMethodNotAllowed)
			done <- true
	}
}

func (r *TaskViewRest) taskViewRestGet(w http.ResponseWriter, req *http.Request, done chan bool) {
	
	done <- true
}

type taskViewRestPostBody struct {
	Title *string `json:"title"`
	Content *string `json:"content"`
}

func (r *TaskViewRest) taskViewRestPost(w http.ResponseWriter, req *http.Request, done chan bool) {
	defer func () {
		done <- true
	}()

	userIdCookie, err := req.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userId, err := strconv.Atoi(userIdCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	body := taskViewRestPostBody{}
	err = decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.Title == nil {
		http.Error(w, "missing field title", http.StatusBadRequest)
		return
	}

	if body.Content == nil {
		http.Error(w, "missing field content", http.StatusBadRequest)
		return
	}

	task, err := r.controller.CreateTaskForUserId(ctx, userId, *body.Title, *body.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskJson, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(taskJson)
}