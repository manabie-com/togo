package app_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	togo "togo/app"
	"togo/helpers"

	"github.com/joho/godotenv"
)

var app = togo.App{}

func TestMain(m *testing.M) {
	if err := godotenv.Load("../app.env"); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	db_params := togo.DB_Params{
		DB_NAME:     os.Getenv("APP_DB_NAME"),
		DB_USERNAME: os.Getenv("APP_DB_USERNAME"),
		DB_PASSWORD: os.Getenv("APP_DB_PASSWORD"),
	}

	app.Initialize(&db_params)

	helpers.EnsureTablesExist(app.DB)
	helpers.CreateInitialUser(app.DB)
	code := m.Run()
	helpers.ClearTables(app.DB)
	os.Exit(code)
}

func makeRequestTo(endpoint, method string, payload []byte) *httptest.ResponseRecorder {
	var request *http.Request
	responseRecorder := httptest.NewRecorder()

	if method == "POST" && payload != nil {
		request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(payload))
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, _ = http.NewRequest(method, endpoint, nil)
	}

	app.Router.ServeHTTP(responseRecorder, request)

	return responseRecorder
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
			response := makeRequestTo(endpoint, method, nil)

			if http.StatusNotFound != response.Code {
				t.Errorf("Expected response code %d. Got %d\n", http.StatusNotFound, response.Code)
			}
		}
	}
}
