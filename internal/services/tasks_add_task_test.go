package services

import (
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddTaskRunner(t *testing.T) {
	t.Run("Scenario 1: Test add task with valid user_id and created_date", func(t *testing.T) {
		testAddTaskWithValidUserIdAndCreatedDate(t, storages.Task{
			Content:     "Do homework",
			UserID:      "firstUser",
			CreatedDate: time.Now().Format("2006-01-02"),
		})
	})

	t.Run("Scenario 2: Test add task with user_id is not exists in table users", func(t *testing.T) {
		testAddTaskWithUserIdIsNotExistInTableUser(t, storages.Task{
			Content:     "Do homework",
			UserID:      "userNotExists",
			CreatedDate: time.Now().Format("2006-01-02"),
		})
	})
}

func testAddTaskWithValidUserIdAndCreatedDate(t *testing.T, input storages.Task) {
	path := config.PathConfig.PathTasks
	response := addTask(path, http.MethodPost, input)
	if response.Code != 200 {
		t.Errorf("Expected status code 200 but the fact is %d", response.Code)
		return
	}

	errorResp := model.ErrorResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &errorResp)
	if err != nil {
		t.Errorf("Error when convert []byte to ErrorResponse. Detail %v", err)
		return
	}

	if errorResp.Error != nil {
		t.Errorf("Expected created task successfully but the fact is error. Detail: %v", errorResp)
		return
	}

	tasks := model.AddTaskResponse{}
	err = json.Unmarshal(response.Body.Bytes(), &tasks)
	if err != nil {
		t.Errorf("Error when convert []byte to AddTaskResponse")
		return
	}

	data := tasks.Data
	if data == nil {
		t.Errorf("Expected data from response but the fact is nil")
		return
	}

	if input.Content != data.Content {
		t.Errorf("Field content in input and field content in task has created  is different. Field content in input is %s and field content in task has created is %s", input.Content, data.Content)
	} else if input.CreatedDate != data.CreatedDate {
		t.Errorf("Field created_date in input and field created_date in is different. Field created_date in input is %s and field created_date in task has created is %s", input.CreatedDate, data.CreatedDate)
	} else if input.UserID != data.UserID {
		t.Errorf("Field user_id in input and field user_id in is different. Field user_id in input is %s and field user_id in task has created is %s", input.UserID, data.UserID)
	}

}

func testAddTaskWithUserIdIsNotExistInTableUser(t *testing.T, input storages.Task) {
	path := config.PathConfig.PathTasks
	response := addTask(path, http.MethodPost, input)
	if response.Code != 404 {
		t.Errorf("Expected status is 404 but the fact is %d", response.Code)
		return
	}

	errorResp := model.ErrorResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &errorResp)
	if err != nil {
		t.Errorf("Error when convert []byte to ErrorResponse. Detail %v", err)
		return
	}

	if errorResp.Error == nil {
		t.Errorf("Expected created task error with error info but the fact is different. Detail: %v", response.Body)
		return
	}
}

func addTask(path string, method string, input storages.Task) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, input.ToIOReader())
	request = request.WithContext(context.WithValue(request.Context(), userAuthKey(0), input.UserID))
	response := httptest.NewRecorder()
	ServiceMockForTest.addTask(response, request)
	return response
}
