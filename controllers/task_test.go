package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// Pass ✅
func TestAdd(t *testing.T) {
	payload := []byte(`{
		"name" : "task name",
		"content" : "task content"
	}`)

	req, _ := http.NewRequest("POST", "/api/tasks/add", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusCreated, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success create task", r.Message)

	if r.Data["name"] != "task name" {
		t.Errorf("Expected task name is 'task name'. Got '%v'", r.Data["name"])
	}
	if r.Data["content"] != "task content" {
		t.Errorf("Expected task content is 'task content'. Got '%v'", r.Data["content"])
	}
}

// Pass ✅
func TestGetTasks(t *testing.T) {
	type tasksResponse struct {
		Status  string
		Message string
		Data    []map[string]interface{}
	}
	// get user id here
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	var tr tasksResponse
	json.Unmarshal(res.Body.Bytes(), &tr)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success", r.Message)
	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if len(r.Data) != 8 {
		t.Errorf("Expected tasks length to be '7'. Got '%v'", len(r.Data))
	}
}

// Pass ✅
func TestGetTask(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tasks/4", nil)
	// auth token
	req.Header.Set("Authorization", token)

	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success", r.Message)

	if r.Data["name"] != "task name" {
		t.Errorf("Expected task name to be 'task name' value. Got '%v'", r.Data["name"])
	}
	if r.Data["content"] != "subtask1 \n subtask2" {
		t.Errorf("Expected task content to be 'subtask1 \n subtask2' value. Got '%v'", r.Data["content"])
	}
}

// Pass ✅
func TestEdit(t *testing.T) {
	payload := []byte(`{
		"name" : "task name",
		"content" : "task content"
	}`)
	req, _ := http.NewRequest("PATCH", "/api/tasks/4", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success update task", r.Message)
	if r.Data["name"] != "task name" {
		t.Errorf("Expected task name is 'task name'. Got '%v'", r.Data["name"])
	}
	if r.Data["content"] != "task content" {
		t.Errorf("Expected task content is 'task content'. Got '%v'", r.Data["content"])
	}
}

// Pass ✅
func TestDelete(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/tasks/3", nil)
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusNoContent, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success delete task", r.Message)

	if r.Data != nil {
		t.Errorf("Expected type of Data to be 'nil' value. Got '%v'", r.Data)
	}
	t.Logf("Response %v", r)
}
