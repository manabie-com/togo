package main_test

import (
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
