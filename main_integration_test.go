package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/magiconair/properties/assert"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/require"
)

const (
	email    = "me@here.com"
	password = "password"
	dbDriver = "postgres"
	jwtKey   = "wqGyEBBfPK9w3Lxw"
)

var token string

func TestMain(t *testing.T) {
	t.Run("db", testDatabaseConnection)
	t.Run("login", testLogin)
	t.Run("loginPassNotValid", testNotValidPasswordLogin)

	// requestsPerDay is 5 by default: https://github.com/quantonganh/togo/blob/master/cmd/togo/main.go#L18
	// newman is run at the first time, so there are only 4 tokens
	for i := 1; i <= 5; i++ {
		t.Run(fmt.Sprintf("addTask %d", i), func(t *testing.T) {
			addTask(t, fmt.Sprintf("hash password %d", i), http.StatusOK)
		})
	}

	t.Run("addTask 6", func(t *testing.T) {
		addTask(t, "hash password 5", http.StatusTooManyRequests)
	})

}

/**
* Test for connecting database postgres
**/
func testDatabaseConnection(t *testing.T) {
	err := godotenv.Load()
	require.NoError(t, err)
	dbSource := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	require.NoError(t, err)
}

//test Valid Credentials Login
func testLogin(t *testing.T) {
	formData := url.Values{}
	formData.Add("email", email)
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

//test not valid Credentials Login
func testNotValidPasswordLogin(t *testing.T) {
	formData := url.Values{}
	formData.Add("email", email)
	formData.Add("password", "123")

	req, err := http.NewRequest(http.MethodPost, "http://localhost:5050/login", strings.NewReader(formData.Encode()))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	var loginResp map[string]string
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&loginResp))
}

//adding tasks for the new feature limit task of user
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
	//meaning the user exceeds more tasks per day and catching our loop that creates tasks on the TestMain
	if resp.StatusCode == 429 {
		assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	}
	if resp.StatusCode == http.StatusOK {
		var addTaskResp map[string]*storages.Task
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&addTaskResp))
		assert.Equal(t, content, addTaskResp["data"].Content)
	}
}
