package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// Pass ✅
func TestSignUp(t *testing.T) {
	payload := []byte(`{
		"name":     "user1",
		"email":    "user10@gmail.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/signup", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Data["email"] != "user10@gmail.com" {
		t.Errorf("Expected type of Data to be 'user10@gmail.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "user1" {
		t.Errorf("Expected type of Data to be 'user1'. Got '%v'", r.Data["name"])
	}
	t.Logf("Response Message: %v", r.Message)
}

// Pass ✅
func TestLogin(t *testing.T) {
	payload := []byte(`{
		"email":    "user1@gmail.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Data["email"] != "user1@gmail.com" {
		t.Errorf("Expected type of Data to be 'user1@gmail.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "user1" {
		t.Errorf("Expected type of Data to be 'user1'. Got '%v'", r.Data["name"])
	}
	t.Logf("Response Message: %v", r.Message)
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
}

// Pass ✅
func TestGetMe(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/users/me", nil)
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)

	if r.Data["email"] != "user1@gmail.com" {
		t.Errorf("Expected field of Data email to be 'test2@example.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "user1" {
		t.Errorf("Expected field of Data name to be 'user1'. Got '%v'", r.Data["name"])
	}
	if r.Data["is_payment"] != false {
		t.Errorf("Expected field of Data is_payment to be 'true'. Got '%v'", r.Data["is_payment"])
	}
	if r.Data["limit_day_tasks"] != 10.0 {
		t.Errorf("Expected field of Data limit_day_tasks to be '10'. Got '%v'", r.Data["limit_day_tasks"])
	}
	t.Logf("Response Message: %v", r.Message)
}

// Pass ✅
func TestUpdateMe(t *testing.T) {
	payload := []byte(`{
		"name": "updated test user",
		"email":    "updatedtest@example.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("PATCH", "/api/users/edit", bytes.NewBuffer(payload))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjExLCJMaW1pdERheVRhc2tzIjoxMH0.0C6oI6DMy6lKnlw8o_VOllwGz7qkGC-955rVDgzEak4"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)

	if r.Data["email"] != "updatedtest@example.com" {
		t.Errorf("Expected field of Data email to be 'updatedtest@example.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "updated test user" {
		t.Errorf("Expected field of Data name to be 'updated test user'. Got '%v'", r.Data["name"])
	}
	t.Logf("Response Message: %v", r.Message)
}

// Pass ✅
func TestDeleteMe(t *testing.T) {
	payload := []byte(`{
		"password": "123456"
	}`)
	req, _ := http.NewRequest("DELETE", "/api/users/delete", bytes.NewBuffer(payload))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, res.Code)

	json.Unmarshal(res.Body.Bytes(), &r)

	if r.Data != nil {
		t.Errorf("Expected field of Data to be 'nil'. Got '%v'", r.Data)
	}
	t.Logf("Response Message: %v", r.Message)
}
