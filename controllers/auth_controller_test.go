package controllers_test

import (
	"TOGO/controllers"
	"TOGO/models"
	"fmt"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Token   string                 `json:"token"`
	Data    map[string]interface{} `json:"data"`
}

func TestSignup(t *testing.T) {
	var jsonStr = []byte(`{"username": "tuanchoitest2", "password": "123456","name":"Nguyen tuan"}`)

	req, err := http.NewRequest("POST", "/user/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Signup())
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if r.Data["username"] != "tuanchoitest2" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["username"], "tuanchoitest2")
	}
	if r.Data["name"] != "Nguyen tuan" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["name"], "Nguyen tuan")
	}

	rq, _ := http.NewRequest("DELETE", fmt.Sprintf("/user/%s", r.Data["id"]), nil)
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA5NTI4OTQsImlkIjoiNjJiZDY0NDRlNTIyYjdhYmQwODY1Mzg3Iiwicm9sZSI6ImFkbWluIn0.BkxuUq7kh_8ebGuZZRrPDbyy1GX3V02rtMWE-0Kvig4"
	rq.Header.Set("Authorization", token)
	rq.Header.Set("Content-Type", "application/json")
	_ = ExcuteRoute(rq)
}

func TestLogin(t *testing.T) {
	var jsonStr = []byte(`{"username": "tuantest", "password": "123456"}`)
	req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Login())
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if r.Data["username"] != "tuantest" {
		t.Errorf("handler returned wrong data: got %v want %v", r.Data["username"], "tuanchoi1")
	}

	if !models.CheckPasswordHash("123456", r.Data["password"].(string)) {
		t.Errorf("handler returned wrong data")
	}
}
