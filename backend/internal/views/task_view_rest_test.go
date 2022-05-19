package views

import (
	"testing"
	"go.uber.org/dig"
	"manabie.com/internal/common"
	"manabie.com/internal/models"
	"manabie.com/internal/repositories"
	"manabie.com/internal/controllers"
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"io"
	"github.com/stretchr/testify/require"
	"bytes"
)

func TestViewRest(t *testing.T) {
	container := dig.New()
	var err error
	err = common.ProvideClockSim(container)
	if err != nil {
		t.Fatal(err)
	}
	err = repositories.ProvidMockeRepositoryFactory(container)
	if err != nil {
		t.Fatal(err)
	}
	
	err = container.Invoke(func(iRepository repositories.RepositoryFactoryI) {
		mock := iRepository.(*repositories.RepositoryFactoryMock)
		mock.InitUsers(100)
	})
	if err != nil {
		t.Fatal(err)
	}

	err = controllers.ProvideTaskController(container)
	if err != nil {
		t.Fatal(err)
	}
	err = ProvideTaskViewRest(container)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("create task fail if no cookie", func (t *testing.T) {
		err := container.Invoke(func (iView *TaskViewRest) {
			server := httptest.NewServer(iView)
			client := server.Client()
			
			requestBody := map[string]interface{} {
				"title": "abc",
				"content": "def",
			}
			requestBodyJson, err := json.Marshal(requestBody)
			req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(requestBodyJson))

			response, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, 400, response.StatusCode, string(bodyBytes))
		})

		if err != nil {
			t.Fatal(err)
		}		
	})

	t.Run("create task fail if missing content", func (t *testing.T) {
		err := container.Invoke(func (iView *TaskViewRest) {
			server := httptest.NewServer(iView)
			client := server.Client()
			
			requestBody := map[string]interface{} {
				"title": "abc",
			}
			cookie := &http.Cookie{
				Name:   "user_id",
				Value:  "1",
				MaxAge: 300,
			}
			requestBodyJson, err := json.Marshal(requestBody)
			req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(requestBodyJson))
			req.AddCookie(cookie)

			response, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, 400, response.StatusCode, string(bodyBytes))
		})

		if err != nil {
			t.Fatal(err)
		}		
	})

	t.Run("create task fail if missing title", func (t *testing.T) {
		err := container.Invoke(func (iView *TaskViewRest) {
			server := httptest.NewServer(iView)
			client := server.Client()
			
			requestBody := map[string]interface{} {
				"content": "abc",
			}
			cookie := &http.Cookie{
				Name:   "user_id",
				Value:  "1",
				MaxAge: 300,
			}
			requestBodyJson, err := json.Marshal(requestBody)
			req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(requestBodyJson))
			req.AddCookie(cookie)

			response, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, 400, response.StatusCode, string(bodyBytes))
		})

		if err != nil {
			t.Fatal(err)
		}		
	})

	t.Run("create task ok", func (t *testing.T) {
		err := container.Invoke(func (iView *TaskViewRest, ) {
			server := httptest.NewServer(iView)
			client := server.Client()

			cookie := &http.Cookie{
				Name:   "user_id",
				Value:  "1",
				MaxAge: 300,
			}
			requestBody := map[string]interface{} {
				"title": "abc",
				"content": "def",
			}
			requestBodyJson, err := json.Marshal(requestBody)
			req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(requestBodyJson))
			if err != nil {
				t.Fatal(err)
			}
			req.AddCookie(cookie)

			response, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, 200, response.StatusCode, string(bodyBytes))

			decoder := json.NewDecoder(bytes.NewBuffer(bodyBytes))
			decoder.DisallowUnknownFields()
			var task models.Task
			err = decoder.Decode(&task)
			if err != nil {
				t.Fatal(err)
			}
			require.NotEqual(t, -1, task.Id)
			require.Equal(t, "abc", task.Title)
			require.Equal(t, "def", task.Content)
			require.True(t, nil == task.Owner)

			err = container.Invoke(func(iRepository repositories.RepositoryFactoryI) {
				mock := iRepository.(*repositories.RepositoryFactoryMock)
				require.Equal(t, 1, len(mock.TaskRepository.Tasks[1]))
				require.Equal(t, task.Id, mock.TaskRepository.Tasks[1][0].Id)
				require.Equal(t, "abc", mock.TaskRepository.Tasks[1][0].Title)
				require.Equal(t, "def", mock.TaskRepository.Tasks[1][0].Content)
			})
			if err != nil {
				t.Fatal(err)
			}
		
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}
