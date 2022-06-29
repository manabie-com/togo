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
		"name":     "test_user",
		"email":    "test_user@gmail.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/signup", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusCreated, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Created Account", r.Message)

	if r.Data["email"] != "test_user@gmail.com" {
		t.Errorf("Expected type of Data to be 'test_user@gmail.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "test_user" {
		t.Errorf("Expected type of Data to be 'test_user'. Got '%v'", r.Data["name"])
	}
}

// Pass ✅
func TestLogin(t *testing.T) {
	payload := []byte(`{
		"email":    "test_user@gmail.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Login Success", r.Message)

	if r.Data["email"] != "test_user@gmail.com" {
		t.Errorf("Expected type of Data to be 'test_user@gmail.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "test_user" {
		t.Errorf("Expected type of Data to be 'test_user'. Got '%v'", r.Data["name"])
	}
}

// Pass ✅
func TestGetMe(t *testing.T) {
	// get token from test user
	token := getToken()
	req, _ := http.NewRequest("GET", "/api/users/me", nil)
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success", r.Message)

	// json.NewDecoder(res.Body).Decode(&m)

	if r.Data["name"] != "test_user" {
		t.Errorf("Expected user name is 'test_user'. Got '%v'", r.Data["name"])
	}
	if r.Data["email"] != "test_user@gmail.com" {
		t.Errorf("Expected user email is 'test_user@gmail.com'. Got '%v'", r.Data["email"])
	}
	if r.Data["is_payment"] != false {
		t.Errorf("Expected user is_payment field is 'true'. Got '%v'", r.Data["is_payment"])
	}
	if r.Data["limit_day_tasks"] != 10.0 {
		t.Errorf("Expected user limit task field is '10'. Got '%v'", r.Data["limit_day_tasks"])
	}
}

// Pass ✅
func TestUpdateMe(t *testing.T) {
	// get token from test user
	token := getToken()
	payload := []byte(`{
		"name": "updated_test_user",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("PATCH", "/api/users/edit", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success update your account!", r.Message)

	// json.NewDecoder(res.Body).Decode(&m)

	if r.Data["email"] != "test_user@gmail.com" {
		t.Errorf("Expected field of Data email to be 'test_user@gmail.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "updated_test_user" {
		t.Errorf("Expected field of Data name to be 'updated_test_user'. Got '%v'", r.Data["name"])
	}
}

// Pass ✅
func TestDeleteMe(t *testing.T) {
	token := getToken()
	payload := []byte(`{
		"password": "123456"
	}`)
	req, _ := http.NewRequest("DELETE", "/api/users/delete", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	json.Unmarshal(res.Body.Bytes(), &r)

	checkResponseCode(t, http.StatusNoContent, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success delete your account!", r.Message)

	if r.Data != nil {
		t.Errorf("Expected field of Data to be 'nil'. Got '%v'", r.Data)
	}
	rollbackUser()
}

func getToken() string {
	type predictResponse struct {
		Status  string
		Message string
		Data    map[string]string
	}
	payload := []byte(`{
		"email":    "test_user@gmail.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	var pr predictResponse
	json.Unmarshal(res.Body.Bytes(), &pr)
	return "Bearer " + pr.Data["token"]
}
