// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/manabie-com/togo/internal/storages"
)

const (
	userID = "firstUser"
	password = "example"
)

var token string

func TestRateLimit(t *testing.T) {
	t.Run("login", testLogin)

	// requestsPerDay is 5 by default: https://github.com/quantonganh/togo/blob/master/cmd/togo/main.go#L18
	// newman is run at the first time, so there are only 4 tokens
	for i := 1; i <= 4; i++ {
		t.Run(fmt.Sprintf("addTask %d", i), func(t *testing.T) {
			addTask(t, fmt.Sprintf("hash password %d", i), http.StatusOK)
		})
	}

	t.Run("addTask 5", func(t *testing.T) {
		addTask(t, "hash password 5", http.StatusTooManyRequests)
	})
}

func testLogin(t *testing.T) {
	formData := url.Values{}
	formData.Add("user_id", userID)
	formData.Add("password", password)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:5050/login", strings.NewReader(formData.Encode()))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var loginResp map[string]string
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&loginResp))
	token = loginResp["data"]
}

func addTask(t *testing.T, content string, statusCode int) {
	data := map[string]string{
		"content": content,
	}
	dataJson, err := json.Marshal(data)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:5050/tasks", bytes.NewBuffer(dataJson))
	require.NoError(t, err)
	req.Header.Add("Authorization", token)

	client := http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, statusCode, resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		var addTaskResp map[string]*storages.Task
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&addTaskResp))
		assert.Equal(t, content, addTaskResp["data"].Content)
	}
}

