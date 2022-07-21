package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	togo "togo/app"
)

var app = togo.App{}

func TestMain(m *testing.M) {
	app.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestApp_ShouldRespondOkToTestRequests(t *testing.T) {
	methods := []string{"GET", "POST"}
	for _, method := range methods {
		response := makeRequestTo("/", method)

		if http.StatusOK != response.Code {
			t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.Code)
		}
	}
}

func TestApp_ShouldHandleNonExistingRoutes(t *testing.T) {
	routes := map[string][]string{
		"POST":   {"/offices", "/are", "/outdated"},
		"GET":    {"/me", "/a", "/pet", "/gopher"},
		"PUT":    {"/some", "/spice", "/in", "/your", "/life"},
		"DELETE": {"/unwated", "/memories", "/from", "/your", "/mind"},
	}

	for method, endpoints := range routes {
		for _, endpoint := range endpoints {
			response := makeRequestTo(endpoint, method)

			if http.StatusNotFound != response.Code {
				t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
			}
		}
	}
}

func makeRequestTo(endpoint, method string) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	request, _ := http.NewRequest(method, endpoint, nil)
	app.Router.ServeHTTP(responseRecorder, request)

	return responseRecorder
}
