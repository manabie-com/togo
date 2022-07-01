package controllers_test

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	// send data to create
	var jsonStr = []byte(`{"name": "test task","content": "test content"}`)
	req, _ := http.NewRequest("POST", "/task", bytes.NewBuffer(jsonStr))

	// send token
	token := tokenMain
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.CreateTask()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusCreated {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}

	if r.Data["name"] != "test task" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "test task")
	}
	if r.Data["content"] != "test content" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["content"], "test content")
	}
}

func TestGetTask(t *testing.T) {
	req, _ := http.NewRequest("GET", "/user-tasks", nil)

	// send token
	token := tokenMain
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.GetTask()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}
}

func TestGetOneTask(t *testing.T) {
	req, _ := http.NewRequest("GET", "/task/62bd689029af520356f8fc0e", nil)

	//send token
	token := tokenMain
	req.Header.Set("Authorization", token)
	res := ExcuteRoute(req)
	var r Response
	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}

	if r.Data["name"] != "Task-tuantest" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "task-tuanchoi1")
	}

}

// user tuantest2
func TestUpdateTask(t *testing.T) {
	var jsonStr = []byte(`{"name": "task-tuantest2", "content": "test update content"}`)
	req, _ := http.NewRequest("PUT", "/task/62bd693829af520356f8fc12", bytes.NewBuffer(jsonStr))

	//send token
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA5NTM5MTcsImlkIjoiNjJiZDY5MzYyOWFmNTIwMzU2ZjhmYzEwIiwicm9sZSI6InVzZXIifQ.-wSRPXY30nxe8P9uzZDyHe8tgXWrb2Vh-LtOU6moS8Y"
	req.Header.Set("Authorization", token)
	res := ExcuteRoute(req)
	var r Response
	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}

	if r.Data["name"] != "task-tuantest2" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "test update task")
	}

	if r.Data["content"] != "test update content" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "test update content")
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	var jsonStr = []byte(`{"status": "completed"}`)
	req, _ := http.NewRequest("PUT", "/task/status/62bd688f29af520356f8fc0c", bytes.NewBuffer(jsonStr))
	//send token
	token := tokenMain
	req.Header.Set("Authorization", token)
	res := ExcuteRoute(req)
	var r Response
	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}

	if r.Data["status"] != "completed" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["status"], "completed")
	}
}

func TestGetTaskDoing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/task-status", nil)
	//send token
	token := tokenMain
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.GetTaskDoing()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong data: got %v want %v", r.Status, http.StatusOK)
	}

	if r.Message != "success" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["message"], "success")
	}
}
