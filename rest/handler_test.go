package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"togo/common/environment"
	"togo/db"
	"togo/models"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

var username = uuid.NewV1().String()

func Router() (*mux.Router, error) {
	environment.Load("../.env")
	db, err := db.Connect()
	restService := Handler(db)

	router := mux.NewRouter()
	router.HandleFunc("/api/user", restService.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", restService.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/user", restService.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/task", restService.CreateTask).Methods("POST")

	return router, err
}

func TestCreateUser(t *testing.T) {

	requestBody, _ := json.Marshal(models.CreateUserRequest{
		Username:       username,
		TaskDailyLimit: 2,
	})

	request, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(requestBody))

	router, errRouter := Router()

	assert.Nil(t, errRouter, "The `errRouter` should be nil")

	response1 := httptest.NewRecorder()

	router.ServeHTTP(response1, request)

	assert.Equal(t, 201, response1.Code, "201 Created is expected")

	response2 := httptest.NewRecorder()

	router.ServeHTTP(response2, request)

	assert.Equal(t, 400, response2.Code, "400 Bad Request is expected")
}

func TestUpdateUser(t *testing.T) {

	requestBody, _ := json.Marshal(models.CreateUserRequest{
		Username:       username,
		TaskDailyLimit: 1,
	})

	request, _ := http.NewRequest("PATCH", "/api/user", bytes.NewBuffer(requestBody))

	router, errRouter := Router()

	assert.Nil(t, errRouter, "The `errRouter` should be nil")

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "200 OK is expected")
}

func TestCreateTask(t *testing.T) {

	description := "Sample description"

	requestBody, _ := json.Marshal(models.CreateTaskRequest{
		Username:    username,
		Title:       "Sample title",
		Description: &description,
	})

	request, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(requestBody))

	router, errRouter := Router()

	assert.Nil(t, errRouter, "The `errRouter` should not be nil")

	response1 := httptest.NewRecorder()

	router.ServeHTTP(response1, request)

	assert.Equal(t, 201, response1.Code, "201 Created is expected")

	response2 := httptest.NewRecorder()

	router.ServeHTTP(response2, request)

	assert.Equal(t, 400, response2.Code, "400 Bad Request is expected")
}

func TestDeleteUser(t *testing.T) {

	requestBody, _ := json.Marshal(models.DeleteUserRequest{
		Username: username,
	})

	request, _ := http.NewRequest("DELETE", "/api/user", bytes.NewBuffer(requestBody))

	response := httptest.NewRecorder()

	router, errRouter := Router()

	assert.Nil(t, errRouter, "The `errRouter` should not be nil")

	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "200 OK is expected")
}
