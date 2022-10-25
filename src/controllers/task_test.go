package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-api/config"
	"todo-api/logger"
	"todo-api/service"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	cfg, _ := config.New()
	s := &service.Service{
		Logger: logger.GetLogger(),
		Config: cfg,
		DB:     config.DBMock(),
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/task", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("UserID", "123456")

	if assert.NoError(t, GetTasks(s)(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestCreateTask(t *testing.T) {
	cfg, _ := config.New()
	s := &service.Service{
		Logger: logger.GetLogger(),
		Config: cfg,
		DB:     config.DBMock(),
	}

	body := `
	{
		"content": "todo"
	}
	`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/task", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("UserID", "123456")

	if assert.NoError(t, CreateTask(s)(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	cfg, _ := config.New()
	s := &service.Service{
		Logger: logger.GetLogger(),
		Config: cfg,
		DB:     config.DBMock(),
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/task/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("UserID", "123456")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteTask(s)(c)) {
		assert.Equal(t, 422, rec.Code)
	}
}
