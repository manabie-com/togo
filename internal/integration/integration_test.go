package pkg

import (
	"bytes"
	"encoding/json"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"
	"lntvan166/togo/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const ADMIN_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6ImFkbWluIn0.ei4kWxPWuJyiIQBok-ojPpwY8CA6NcFw-APrjOuI_rk"
const ADMIN = "admin"

const (
	// Config is the global config
	DATABASE_URL string = ""
	PORT         string = ""
	FREE_LIMIT   int    = 0
	VIP_LIMIT    int    = 0
	HOST         string = "http://localhost:8080"
)

var Handler *delivery.Handler

func setupIntegrationTest() *delivery.Handler {
	config.Load("testing")

	db := repository.Connect()

	return delivery.NewHandler(db)
}

func teardownIntegrationTest() {
}

func TestRegisterIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", HOST+"/auth/register", strings.NewReader(`
	{
		"username": "test_integration",
		"password":"admin"
	}`))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	route := mux.NewRouter()
	route.HandleFunc("/auth/register", handler.Register).Methods("POST")
	route.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	userID, _ := handler.UserDelivery.UserUsecase.GetUserIDByUsername("test_integration")
	handler.UserDelivery.UserUsecase.DeleteUserByID(userID)
}

func TestLoginIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", HOST+"/auth/login", strings.NewReader(`
	{
		"username": "admin",
		"password":"admin"
	}`))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	route := mux.NewRouter()
	route.HandleFunc("/auth/login", handler.Login).Methods("POST")
	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	// get token from body
	var element map[string]interface{}
	json.Unmarshal([]byte(body), &element)

	token := element["token"].(string)

	assert.NotEmpty(t, token)
}

func TestGetAllUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("GET", HOST+"/user", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/user", handler.GetAllUsers).Methods("GET")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}

func TestGetUserIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("GET", HOST+"/user/1", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/user/{id}", handler.GetUser).Methods("GET")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}

func TestGetAllTasksIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("GET", HOST+"/task", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/task", handler.GetAllTaskOfUser).Methods("GET")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}

func TestGetTaskIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("GET", HOST+"/task/1", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/task/{id}", handler.GetTaskByID).Methods("GET")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}

func TestCreateTaskIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", HOST+"/task", strings.NewReader(`
	{
		"title": "test_integration",
		"description": "test_integration",
	}`))
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/task", handler.CreateTask).Methods("POST")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusCreated, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	// get task id from body
	var element map[string]interface{}
	json.Unmarshal([]byte(body), &element)

	taskID := element["taskID"].(float64)

	assert.NotEmpty(t, taskID)

	handler.TaskDelivery.TaskUsecase.DeleteTask(int(taskID), ADMIN)
}

func TestCompleteTaskIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", HOST+"/task/1", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/task/{id}", handler.CompleteTask).Methods("POST")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}

func TestDeleteTaskIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	task, _ := handler.TaskDelivery.TaskUsecase.GetTaskByID(1, ADMIN) // prepare for rollback

	r, err := http.NewRequest("DELETE", HOST+"/task/1", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/task/{id}", handler.DeleteTask).Methods("DELETE")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)

	// rollback
	handler.TaskDelivery.TaskUsecase.RollbackFromDelete(task)
}

func TestGetPlanIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("GET", HOST+"/plan", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/plan", handler.GetPlan).Methods("GET")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}

func TestUpgradePlanIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", HOST+"/plan/upgrade/2", strings.NewReader(`
	{
		"plan": "premium"
	}`))
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/plan/upgrade/{id}", handler.UpgradePlan).Methods("POST")
	route.Use(middleware.Authorization)

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)

	handler.UserDelivery.UserUsecase.UpgradePlan(2, "free", 10) // downgrade plan
}
