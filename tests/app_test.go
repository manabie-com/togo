package main_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"togo/app"
)

func TestApp_TypeShouldExist(t *testing.T) {
	appVar := app.App{}
	appVarType := reflect.TypeOf(appVar)

	if appVarType.Name() != "App" && appVarType.PkgPath() == "" {
		t.Error("Type does not exist")
	}
}

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
