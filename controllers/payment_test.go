package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"
)

// Pass ✅
func TestPayment(t *testing.T) {
	signup()
	token := getToken()
	req, _ := http.NewRequest("POST", "/api/payments", nil)
	req.Header.Set("Authorization", token)
	res := executeRequest(req)

	json.Unmarshal(res.Body.Bytes(), &r)
	checkResponseCode(t, http.StatusOK, res.Code)
	checkResponseStatus(t, "Success", r.Status)
	checkResponseMessage(t, "Success upgrade Premium account. Please login again to try new upgrade", r.Message)

	if r.Data["name"] != "test_user" {
		t.Errorf("Expected user name is 'test_user'. Got '%v'", r.Data["name"])
	}
	if r.Data["email"] != "test_user@gmail.com" {
		t.Errorf("Expected user email is 'test_user@gmail.com'. Got '%v'", r.Data["email"])
	}
	if r.Data["is_payment"] != true {
		t.Errorf("Expected user is_payment field is 'true'. Got '%v'", r.Data["is_payment"])
	}
	if r.Data["limit_day_tasks"] != 20.0 {
		t.Errorf("Expected user limit task field is '20'. Got '%v'", r.Data["limit_day_tasks"])
	}
	// rollback before Payments
	rollbackUser()
}
