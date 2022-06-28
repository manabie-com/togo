package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// Pass ✅
func TestGetTasks(t *testing.T) {
	type tasksRes struct {
		Status  string
		Message string
		Data    []map[string]interface{}
	}
	// get user id here
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)
	var r tasksRes
	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if len(r.Data) != 8 {
		t.Errorf("Expected tasks length to be '7'. Got '%v'", len(r.Data))
	}

	for _, val := range r.Data {
		if val["id"] == 4 {
			if val["name"] != "task name" {
				t.Errorf("Expected type of Data to be 'task name' value. Got '%v'", val["name"])
			}
			if val["content"] != "subtask1 \n subtask2" {
				t.Errorf("Expected type of Data to be 'subtask1 \n subtask2' value. Got '%v'", val["content"])
			}
		}
	}
}

// Pass ✅
func TestGetTask(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tasks/4", nil)
	// auth token
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)

	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusOK, res.Code)
	if r.Data["name"] != "task name" {
		t.Errorf("Expected type of Data to be 'task name' value. Got '%v'", r.Data["name"])
	}
	if r.Data["content"] != "subtask1 \n subtask2" {
		t.Errorf("Expected type of Data to be 'subtask1 \n subtask2' value. Got '%v'", r.Data["content"])
	}
}

// Pass ✅
func TestAdd(t *testing.T) {
	payload := []byte(`{
		"name" : "task name",
		"content" : "task content"
	}`)

	req, _ := http.NewRequest("POST", "/api/tasks/add", bytes.NewBuffer(payload))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusCreated, res.Code)

	if r.Data["name"] != "task name" {
		t.Errorf("Expected task name is 'task name'. Got '%v'", r.Data["name"])
	}
	if r.Data["content"] != "task content" {
		t.Errorf("Expected task content is 'task content'. Got '%v'", r.Data["content"])
	}
}

// Pass ✅
func TestEdit(t *testing.T) {
	payload := []byte(`{
		"name" : "task name",
		"content" : "task content"
	}`)
	req, _ := http.NewRequest("PATCH", "/api/tasks/4", bytes.NewBuffer(payload))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusOK, res.Code)
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
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusNoContent, res.Code)
	if r.Data != nil {
		t.Errorf("Expected type of Data to be 'nil' value. Got '%v'", r.Data)
	}
	t.Logf("Response %v", r)
}
