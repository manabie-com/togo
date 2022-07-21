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

func TestApp_ShouldHandleNonExistingRoutes(t *testing.T) {
	app := app.App{}
	app.Initialize()
	routes := map[string][]string{
		"POST":   {"/offices", "/are", "/outdated"},
		"GET":    {"/me", "/a", "/pet", "/gopher"},
		"PUT":    {"/some", "/spice", "/in", "/your", "/life"},
		"DELETE": {"/unwated", "/memories", "/from", "/your", "/mind"},
	}

	for method, endpoints := range routes {
		for _, endpoint := range endpoints {
			request, _ := http.NewRequest(method, endpoint, nil)
			response := httptest.NewRecorder()
			app.Router.ServeHTTP(response, request)

			if http.StatusNotFound != response.Code {
				t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
			}
		}
	}
}
