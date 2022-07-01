package controllers_test

import (
	"TOGO/controllers"
	"TOGO/middleware"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMe(t *testing.T) {
	req, _ := http.NewRequest("GET", "/me", nil)
	token := tokenMain
	req.Header.Set("Authorization", token)

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.AuthMiddleware(controllers.GetMe()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	//check satatus code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//check user
	if r.Data["username"] != "tuantest" {
		t.Errorf("handler returned wrong value: got %v want %v",
			r.Data["username"], "tuantest")
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/62bd682629af520356f8fc0a", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := ExcuteRoute(req)
	var r Response
	json.Unmarshal(res.Body.Bytes(), &r)

	if status := r.Status; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if r.Data["username"] != "tuantest" {
		t.Errorf("handler returned wrong  value: got %v want %v",
			r.Data["username"], "tuanchoi1")
	}

}

func TestUpdateMe(t *testing.T) {
	var jsonStr = []byte(`{"name": "Test tuandz", "password": "123456"}`)
	req, _ := http.NewRequest("PUT", "/user", bytes.NewBuffer(jsonStr))
	token := tokenMain
	req.Header.Set("Authorization", token)

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.UpdateMe()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Status, http.StatusOK)
	}

	if r.Data["name"] != "Test tuandz" {
		t.Errorf("handler returned wrong value: got %v want %v",
			r.Data["name"], "Test tuandz")
	}
}

func TestDeleteUser(t *testing.T) {
	rq, _ := http.NewRequest("POST", "/user/login", nil)
	handler := http.HandlerFunc(CreateTestUser())
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, rq)
	token := tokenAdmin
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user/%s", NewId.Hex()), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	res := ExcuteRoute(req)
	var r Response
	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Status, http.StatusOK)
	}
	if r.Message != "success" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Message, "success")
	}
}

func TestGetAllUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	token := tokenAdmin
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(middleware.AuthMiddleware(controllers.GetAllUser()))
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Status, http.StatusOK)
	}
}

func TestUpdateLimit(t *testing.T) {
	req, err := http.NewRequest("PUT", "/limit", nil)
	if err != nil {
		t.Fatal(err)
	}
	token := tokenMain
	req.Header.Set("Authorization", token)

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.Handler(middleware.AuthMiddleware(controllers.UpdateLimit()))
	handler.ServeHTTP(rr, req)
	var r Response
	json.Unmarshal(rr.Body.Bytes(), &r)

	if r.Status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			r.Status, http.StatusOK)
	}

	if r.Data["limit"].(float64) != 100 {
		t.Errorf("handler returned wrong: got %v want %v",
			r.Data["limit"], "100")
	}
}
