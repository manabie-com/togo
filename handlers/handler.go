package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/SVincentTran/togo/errors"
	"github.com/SVincentTran/togo/forms"
	"github.com/SVincentTran/togo/services"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *services.Service
}

// Http Handler Function for recording to do task
func (h Handler) TodoTasksHanlder(w http.ResponseWriter, req *http.Request) error {
	// Parse request body to form
	todoTaskReq, err := ParseRequestBody(req)
	if err != nil {
		log.Panicf("Get error while parsing request's body %s", err)
	}

	// Get User Id from query vars
	userId, err := strconv.ParseInt(mux.Vars(req)["userId"], 10, 32)
	if err != nil || userId <= 0 {
		return errors.GetError(errors.BadRequestContext, errors.BadRequestMessage, errors.TodoTaskRequestInvalid)
	}
	todoTaskReq.UserId = int(userId)

	// Validate request
	if err := todoTaskReq.Validate(); err != nil {
		return err
	}

	log.Printf("Receiving a request for user: %d, processing to record to do task...", userId)

	// Call the recording service
	err = h.service.RecordTodoTasks(*todoTaskReq)
	if err != nil {
		return err
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := CreatedResponse()
	return json.NewEncoder(w).Encode(response)
}

// Parse json request body to struct
func ParseRequestBody(req *http.Request) (*forms.TodoTaskRequest, error) {
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	todoTaskReq := &forms.TodoTaskRequest{}

	if err != nil {
		return todoTaskReq, err
	}

	if err := json.Unmarshal(body, todoTaskReq); err != nil {
		return todoTaskReq, err
	}

	return todoTaskReq, nil
}

// Create Http response
func CreatedResponse() *forms.BaseResponse {
	response := &forms.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Todo task recorded",
	}

	return response
}

func New() *Handler {
	service := services.New()
	return &Handler{
		service: service,
	}
}
