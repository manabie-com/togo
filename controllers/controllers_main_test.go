package controllers_test

import (
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/manabie-com/togo/app"
)

type response struct {
	Status  string
	Message string
	Data    map[string]interface{}
}

var (
	a     app.App
	r     response
	token string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjUyLCJMaW1pdERheVRhc2tzIjoxMH0.PcgDnM8-LY0oyNiSfYRyQVIMryU1TfqWgEgq6geqEXc"
)

func TestMain(m *testing.M) {
	a = app.App{}
	a.Init()
	code := m.Run()
	os.Exit(code)
}

func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, r)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseStatus(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response Status %s. Got %s\n", expected, actual)
	}
}

func checkResponseMessage(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response Message %s. Got %s\n", expected, actual)
	}
}

func rollbackUser() error {
	_, err := a.DB.Exec(`DELETE FROM users WHERE email = $1`, "test_user@gmail.com")
	if err != nil {
		return err
	}
	return nil
}

func rollbackPayment(email interface{}) error {
	_, err := a.DB.Exec(`UPDATE users SET is_payment = $1, limit_day_tasks = $2 WHERE email = $3`, false, 10, email)
	if err != nil {
		return err
	}
	return nil
}
