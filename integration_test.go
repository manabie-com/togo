package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	_ "togo/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

// Test post task happy
func TestPostTaskHappy(t *testing.T) {
	body := `{
		"summary": 		"Todo task 2022-05-27",
		"description": 	"do something",
		"assignee": 	"IYadf5AYZYZByyTTl1f5QqxOGx13",
		"taskDate": 	"2022-05-27"
	}`
	r, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestPostTaskHappy", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// Test post task bad request
func TestPostTaskBadRequest(t *testing.T) {
	body := `{
		"description": 	"do something",
		"assignee": 	"IYadf5AYZYZByyTTl1f5QqxOGx13",
		"taskDate": 	"2022-05-27"
	}`
	r, _ := http.NewRequest("POST", "/v1/tasks", strings.NewReader(body))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Trace("testing", "TestPostTaskBadRequest", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Endpoint\n", t, func() {
		Convey("Status Code Should Be 400", func() {
			So(w.Code, ShouldEqual, 400)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// Test get health check
func TestGetHealthCheck(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestHealthCheck", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}