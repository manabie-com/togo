package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var SuccessCode int = 0
var MaxLimitCode int = 1
var host string = "http://localhost:8000"

// Error represents our Error in JSON
type Error struct {
	ErrorCode int    `json:"err_code"`
	ErrorDesc string `json:"err_desc"`
}

type Task struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

// CreateTaskResult is a JSON object needed to create a task
type CreateTaskRequest struct {
	Name string `json:"name"`
}

// CreateTaskResult is a JSON object we'll return in our API
type CreateTaskResult struct {
	Error `json:"error"`
	*Task `json:"task,omitempty"`
}

func CreateTask(userID int, taskName string) (int, *CreateTaskResult, error) {
	createTaskReq := CreateTaskRequest{
		Name: taskName,
	}
	reqBody, _ := json.Marshal(createTaskReq)
	url := host + "/users/" + strconv.Itoa(userID) + "/tasks"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Failed to create task with error %v", err)
		return 0, nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Got non 200 status code %d", resp.StatusCode)
		return resp.StatusCode, nil, nil
	}
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read resp body with error %v", err)
		return 0, nil, err
	}
	log.Printf("Got response body '%s'", body)
	createTaskRes := CreateTaskResult{}
	err = json.Unmarshal(body, &createTaskRes)
	if err != nil {
		log.Printf("Failed to read unmarshal response with error %v", err)
		return 0, nil, err
	}
	return resp.StatusCode, &createTaskRes, nil
}
