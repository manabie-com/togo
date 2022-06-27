package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSignUp(t *testing.T) {
	payload := []byte(`{
		"name":     "test user",
		"email":    "test21@example.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/signup", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if r.Message != "Created Account" {
		t.Errorf("Expected Message field to be 'Created Account'. Got '%v'", r.Message)
	}

	if r.Data["email"] != "test21@example.com" {
		t.Errorf("Expected type of Data to be 'test21@example.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "test user" {
		t.Errorf("Expected type of Data to be 'test user'. Got '%v'", r.Data["name"])
	}
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
}

func TestLogin(t *testing.T) {
	payload := []byte(`{
		"email":    "test2@example.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/users/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if r.Message != "Login Success" {
		t.Errorf("Expected Message field to be 'Created Account'. Got '%v'", r.Message)
	}

	if r.Data["email"] != "test2@example.com" {
		t.Errorf("Expected type of Data to be 'test2@example.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "test user1" {
		t.Errorf("Expected type of Data to be 'test user1'. Got '%v'", r.Data["name"])
	}
	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
}

func TestGetMe(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/users/me", nil)
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIzLCJMaW1pdERheVRhc2tzIjoxMH0.L0uyMcBzvzzK6c1hisfXuL0Cp-Gbwu6qiSMVq0Ojwaw"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if r.Message != "Success" {
		t.Errorf("Expected Message field to be 'Success'. Got '%v'", r.Message)
	}

	if r.Data["email"] != "test2@example.com" {
		t.Errorf("Expected field of Data email to be 'test2@example.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "test user1" {
		t.Errorf("Expected field of Data name to be 'test user'. Got '%v'", r.Data["name"])
	}
	if r.Data["is_payment"] != false {
		t.Errorf("Expected field of Data is_payment to be 'true'. Got '%v'", r.Data["is_payment"])
	}
	if r.Data["limit_day_tasks"] != 10.0 {
		t.Errorf("Expected field of Data limit_day_tasks to be '10'. Got '%v'", r.Data["limit_day_tasks"])
	}
}
func TestUpdateMe(t *testing.T) {
	payload := []byte(`{
		"name": "updated test user",
		"email":    "updatedtest@example.com",
		"password": "123456"
	}`)
	req, _ := http.NewRequest("PATCH", "/api/users/edit", bytes.NewBuffer(payload))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIwLCJMaW1pdERheVRhc2tzIjoxMH0.8UHK64DHDcly-fSB22HdsVAoh5y_P_YhN-eX_SZoO_w"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if r.Message != "Success update your account!" {
		t.Errorf("Expected Message field to be 'Success'. Got '%v'", r.Message)
	}

	if r.Data["email"] != "updatedtest@example.com" {
		t.Errorf("Expected field of Data email to be 'updatedtest@example.com'. Got '%v'", r.Data["email"])
	}

	if r.Data["name"] != "updated test user" {
		t.Errorf("Expected field of Data name to be 'updated test user'. Got '%v'", r.Data["name"])
	}
}

func TestDeleteMe(t *testing.T) {
	payload := []byte(`{
		"password": "123456"
	}`)
	req, _ := http.NewRequest("DELETE", "/api/users/delete", bytes.NewBuffer(payload))
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIzLCJMaW1pdERheVRhc2tzIjoxMH0.L0uyMcBzvzzK6c1hisfXuL0Cp-Gbwu6qiSMVq0Ojwaw"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, res.Code)

	var r response
	json.Unmarshal(res.Body.Bytes(), &r)

	// json.NewDecoder(res.Body).Decode(&m)
	if r.Status != "Success" {
		t.Errorf("Expected Status field to be 'Success' but. Got '%v'", r.Status)
	}

	if r.Message != "Success delete your account!" {
		t.Errorf("Expected Message field to be 'Success'. Got '%v'", r.Message)
	}

	if r.Data != nil {
		t.Errorf("Expected field of Data to be 'nil'. Got '%v'", r.Data)
	}
}
