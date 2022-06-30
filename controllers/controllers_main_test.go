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
	a app.App
	r response
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

func rollbackTask() error {
	_, err := a.DB.Exec(`DELETE FROM tasks WHERE user_id = (SELECT users.id FROM users where email = $1)`, "test_user@gmail.com")
	if err != nil {
		return err
	}
	return nil
}
