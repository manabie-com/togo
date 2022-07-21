package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"togo/app"
)

func TestApp_ShouldRespondOkToTestRequests(t *testing.T) {
	app := app.App{}
	app.Initialize()
	methods := []string{"GET", "POST"}
	for _, method := range methods {
		request, _ := http.NewRequest(method, "/", nil)
		response := httptest.NewRecorder()
		app.Router.ServeHTTP(response, request)

		if http.StatusOK != response.Code {
			t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
		}
	}
}
