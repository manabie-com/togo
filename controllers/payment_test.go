package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"
)

// Pass ✅
func TestPayment(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/payments", nil)
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsIkxpbWl0RGF5VGFza3MiOjEwfQ.p71B7Pd2SydLxeVm5LqZIQXgljHAOJN9z5we4wlW2lA"
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusOK, res.Code)
	if r.Data["name"] != "user1" {
		t.Errorf("Expected user name is 'user1'. Got '%v'", r.Data["name"])
	}
	if r.Data["email"] != "user1@gmail.com" {
		t.Errorf("Expected user email is 'user1@gmail.com'. Got '%v'", r.Data["email"])
	}
	if r.Data["is_payment"] != true {
		t.Errorf("Expected user is_payment field is 'true'. Got '%v'", r.Data["is_payment"])
	}
	if r.Data["limit_day_tasks"] != 20.0 {
		t.Errorf("Expected user limit task field is '20'. Got '%v'", r.Data["limit_day_tasks"])
	}
}
