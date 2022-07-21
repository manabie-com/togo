package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"togo/app"
)

func TestApp_ShouldRespondToBasicGetRequest(t *testing.T) {
	app := app.App{}
	app.Initialize()
	request, _ := http.NewRequest("GET", "/", nil)
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, request)

	if http.StatusOK != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}
}

func TestApp_ShouldRespondToBasicPostRequest(t *testing.T) {
	app := app.App{}
	app.Initialize()
	request, _ := http.NewRequest("POST", "/", nil)
	responseRecorder := httptest.NewRecorder()
	app.Router.ServeHTTP(responseRecorder, request)

	if http.StatusOK != responseRecorder.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, responseRecorder.Code)
	}
}
