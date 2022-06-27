package controllers_test

import (
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/manabie-com/togo/app"
)

var a app.App

type response struct {
	Status  string
	Message string
	Data    map[string]interface{}
}

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
