package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetListRunner(t *testing.T) {
	t.Run("Scenario 1: Test get list task with valid user_id and created_date", func(t *testing.T) {
		input := storages.Task{
			UserID:      "firstUser",
			CreatedDate: "2020-06-29",
		}
		testGetListTaskWithValidUserIdAndCreatedDate(t, input, fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", input.CreatedDate))
	})

	t.Run("Scenario 2: Test get list task with invalid user_id and valid created_date", func(t *testing.T) {
		input := storages.Task{
			UserID:      "invalidUser",
			CreatedDate: "2020-06-29",
		}
		testGetListTaskWithInvalidInfo(t, input, fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", input.CreatedDate))
	})

	t.Run("Scenario 3: Test get list task with valid user_id and invalid created_date", func(t *testing.T) {
		input := storages.Task{
			UserID:      "firstUser",
			CreatedDate: "abcxyz",
		}
		testGetListTaskWithInvalidInfo(t, input, fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", input.CreatedDate))
	})

	t.Run("Scenario 4: Test get list task with invalid user_id and invalid created_date", func(t *testing.T) {
		input := storages.Task{
			UserID:      "invalidUser",
			CreatedDate: "abcxyz",
		}
		testGetListTaskWithInvalidInfo(t, input, fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", input.CreatedDate))
	})

	t.Run("Scenario 5: Test get list task with empty created_date param in url", func(t *testing.T) {
		testGetListTaskWithInvalidInfo(t, storages.Task{}, config.PathConfig.PathTasks)
	})

}

func testGetListTaskWithValidUserIdAndCreatedDate(t *testing.T, input storages.Task, getTasksPath string) {
	response := getListTask(getTasksPath, http.MethodGet, input)
	if response.Code != 200 {
		t.Errorf("Expected status code 200 but the fact is %d", response.Code)
	} else {
		tasks := model.GetTaskResponse{}
		err := json.Unmarshal(response.Body.Bytes(), &tasks)
		if err != nil {
			t.Error("Error when convert response body to data")
			return
		}
		if tasks.Data == nil || len(tasks.Data) == 0 {
			t.Error("Expected response data but the fact is null")
			return
		}
	}
}

func testGetListTaskWithInvalidInfo(t *testing.T, input storages.Task, getTasksPath string) {
	response := getListTask(getTasksPath, http.MethodGet, input)
	if response.Code != 200 {
		t.Errorf("Expected status code 200 but the fact is %d", response.Code)
		return
	}

	tasks := model.GetTaskResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &tasks)
	if err != nil {
		t.Error("Error when convert response body to data")
		return
	}

	if !(tasks.Data == nil || len(tasks.Data) == 0) {
		t.Errorf("Expected field data in body is null(nil), but the fact is different")
	}

}

func getListTask(path string, method string, input storages.Task) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, nil)
	request = request.WithContext(context.WithValue(request.Context(), userAuthKey(0), input.UserID))
	response := httptest.NewRecorder()
	ServiceMockForTest.listTasks(response, request)
	return response
}
