package test

import (
	"github.com/jinzhu/configor"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/controller"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskControllerSuccess(t *testing.T) {
	appCfg := &config.Config{}
	err1 := configor.Load(appCfg, "../config.yml")
	if err1 != nil {
		log.Fatal(err1)
	}

	db, err := config.GetPostgersDB(appCfg.DB.Host, appCfg.DB.Port, appCfg.DB.User, appCfg.DB.Password, appCfg.DB.Name)

	task :=controller.NewTaskController(db)
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(task.ListTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}