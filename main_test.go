package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Integration_BadRequest(t *testing.T) {
	router := setUpRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/task/record", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"message\":\"invalid request body\"}", w.Body.String())
}

func Test_Integration_UserNotExist(t *testing.T) {
	router := setUpRoute()

	w := httptest.NewRecorder()
	values := map[string]string{"user_id": "not exist user id", "task": "todo"}
	jsonData, _ := json.Marshal(values)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/task/record",  bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"message\":\"record user's task error: user id does not exist: not exist user id\"}", w.Body.String())
}