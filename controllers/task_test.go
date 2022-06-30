package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// Pass ✅
func TestAdd(t *testing.T) {
	signup()
	token := getToken()
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
		Data    []map[string]string
	}
	signup()
	token := getToken()
	// get user id here
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	tr := tasksResponse{}
	json.Unmarshal(res.Body.Bytes(), &tr)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", tr.Status)
	checkResponseMessage(t, "Success", tr.Message)
	// json.NewDecoder(res.Body).Decode(&m)

	for _, v := range tr.Data {
		if v["name"] != "task name" {
			t.Errorf("Expected task name is 'task name'. Got '%v'", v["name"])
		}
		if v["content"] != "task content" {
			t.Errorf("Expected task content is 'task content'. Got '%v'", v["content"])
		}
	}
}

// Pass ✅
func TestGetTask(t *testing.T) {
	signup()
	token := getToken()
	id := getIdFromCreatedTask(token)
	req, _ := http.NewRequest("GET", "/api/tasks/"+id, nil)
	// auth token
	req.Header.Set("Authorization", token)

	res := executeRequest(req)
	re := response{}
	json.Unmarshal(res.Body.Bytes(), &re)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", re.Status)
	checkResponseMessage(t, "Success", re.Message)

	if re.Data["name"] != "task name" {
		t.Errorf("Expected task name to be 'task name' value. Got '%v'", re.Data["name"])
	}
	if re.Data["content"] != "task content" {
		t.Errorf("Expected task content to be 'task content' value'. Got '%v'", re.Data["content"])
	}
}

// Pass ✅
func TestEdit(t *testing.T) {
	signup()
	token := getToken()
	id := getIdFromCreatedTask(token)
	payload := []byte(`{
		"name" : "task name",
		"content" : "task content"
	}`)
	req, _ := http.NewRequest("PATCH", "/api/tasks/"+id, bytes.NewBuffer(payload))
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
	signup()
	token := getToken()
	id := getIdFromCreatedTask(token)
	req, _ := http.NewRequest("DELETE", "/api/tasks/"+id, nil)
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusNoContent, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success delete task", r.Message)

	if r.Data != nil {
		t.Errorf("Expected type of Data to be 'nil' value. Got '%v'", r.Data)
	}
	rollbackTask()
	rollbackUser()
}

func signup() {
	payload := []byte(`{
		"name":     "test_user",
		"email":    "test_user@gmail.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/signup", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)
	if r.Status == "Failure" {
		return
	}
}

func getIdFromCreatedTask(token string) string {
	type createTaskRes struct {
		Status  string
		Message string
		Data    map[string]string
	}
	payload := []byte(`{
		"name" : "task name",
		"content" : "task content"
	}`)
	req, _ := http.NewRequest("POST", "/api/tasks/add", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	r := createTaskRes{}
	json.Unmarshal(res.Body.Bytes(), &r)
	return r.Data["id"]
}
