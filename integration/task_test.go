package integration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	serverUrl = "http://localhost:8001"
)

func TestGetTask(t *testing.T) {
	t.Run("get task no error", func(t *testing.T) {
		url := fmt.Sprintf("%s%s", serverUrl, "/task?create_task=2022-03-14")
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)
		req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzY1MTQ4OX0.NOg7YJCEaAuzty7MYFjNzTeHSe2W5-8DoLNsBYGpNMQ")
		res, err := client.Do(req)
		assert.NoError(t, err)
		byteBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"status":200,"message":"Success","data":null,"error":null}`, strings.Trim(string(byteBody), "\n"))
		assert.EqualValues(t, http.StatusOK, res.StatusCode)
		res.Body.Close()

	})
	t.Run("get task success", func(t *testing.T) {
		url := fmt.Sprintf("%s%s", serverUrl, "/task?create_date=2022-03-14")
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)
		req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiZmlyc3RVc2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTY0NzY1MTQ4OX0.NOg7YJCEaAuzty7MYFjNzTeHSe2W5-8DoLNsBYGpNMQ")
		res, err := client.Do(req)
		assert.NoError(t, err)
		byteBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"status":200,"message":"Success","data":[{"id":"cca3feb5-f7b5-4f44-951d-4bd418dcad89","content":"new content task","user_id":"firstUser","created_date":"2022-03-14"},{"id":"2bd29cba-e15f-4e2a-af25-891d772fc9df","content":"new content task","user_id":"firstUser","created_date":"2022-03-14"},{"id":"a6483000-4af6-4745-bb7a-57c0225420ae","content":"new content task","user_id":"firstUser","created_date":"2022-03-14"},{"id":"b249a2ca-1c14-446d-98e7-5f65a11d9116","content":"new content task","user_id":"firstUser","created_date":"2022-03-14"},{"id":"2b09706f-5150-4f73-afd8-8ee838ce9024","content":"new content task","user_id":"firstUser","created_date":"2022-03-14"}],"error":null}`, strings.Trim(string(byteBody), "\n"))
		assert.EqualValues(t, http.StatusOK, res.StatusCode)
		res.Body.Close()

	})
	t.Run("get task error Authenticate", func(t *testing.T) {
		url := fmt.Sprintf("%s%s", serverUrl, "/task?create_date=2022-03-14")
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)
		req.Header.Add("Authorization", "invalid token")
		res, err := client.Do(req)
		assert.NoError(t, err)
		byteBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, `{"status":400,"message":"Error :Unauthenticated","data":null,"error":"Unauthenticated"}`, strings.Trim(string(byteBody), "\n"))
		assert.EqualValues(t, http.StatusOK, res.StatusCode)
		res.Body.Close()

	})
}
