package integration_test

import (
	"encoding/json"
	"fmt"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/model"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	t.Run("Scenario 1: Test login with valid user info and get task from this user (user has tasks in database with created_date is 2020-06-29) => Expected login success and receive list task", func(t *testing.T) {
		testLoginAndGetTasksWithUserHasTaskAtCreatedDate(t, "firstUser", "example", "2020-06-29")
	})

	t.Run("Scenario 2: Test login with valid user info and get task from this user (user don't have task with created_date is 2021-03-08)", func(t *testing.T) {
		testLoginAndGetTasksWithUserDontHaveTaskAtCreatedDate(t, "firstUser", "example", "2021-03-08")
	})

	t.Run("Scenario 3: Test get task with valid token from user has task (with created_date is 2020-06-29)", func(t *testing.T) {
		testGetTasksWithValidTokenFromUserHasTaskAtCreatedDate(t, "firstUser", "2020-06-29")
	})

	t.Run("Scenario 4: Test get task with valid token from user don't have task (with created_date is 2021-03-08", func(t *testing.T) {
		testGetTaskWithValidTokenFromUserDontHaveTasksAtCreatedDate(t, "firstUser", "2021-03-08")
	})

	t.Run("Scenario 5: Test get task with valid token from user has task (with created_date is 2020-06-29), but token is expire", func(t *testing.T) {
		testGetTasksWithValidTokenButExpireFromUserHasTask(t, "firstUser", "2020-06-29")
	})

	t.Run("Scenario 6: Test get task with token from user is not exist in database", func(t *testing.T) {
		testGetTaskWithValidTokenFromUserIsNotExistsInDatabase(t, "userNotExist", "2020-06-29")
	})

	t.Run("Scenario 7: Test login and add task", func(t *testing.T) {
		testLoginAndAddTask(t, "firstUser", "example", storages.Task{
			Content: "Do homework",
		})
	})

	t.Run("Scenario 8: Test login and add task with user has reached to limit of task in day", func(t *testing.T) {
		listTask := make([]storages.Task, 0, 5)
		for i := 1; i <= 6; i++ {
			listTask = append(listTask, storages.Task{Content: "Do homework " + strconv.Itoa(i)})
		}
		testLoginAndAddTaskWithUserHasReachedToLimitOfTaskInDay(t, "secondUser", "example", listTask)
	})

	t.Run("Scenario 9: Test add task with valid token", func(t *testing.T) {
		userId := "firstUser"
		token, _ := services.CreateTokenForTest(userId, time.Now().Add(15*time.Minute).Unix())
		testAddTask(t, token, userId, storages.Task{
			Content: "Do homework",
		})
	})

	t.Run("Scenario 10: Test add task with valid token but expire", func(t *testing.T) {
		testAddTaskWithValidTokenButExpire(t, storages.Task{
			UserID:  "firstUser",
			Content: "Do homework",
		})
	})

	t.Run("Scenario 11: Test add task with invalid token", func(t *testing.T) {
		testAddTaskWithInvalidToken(t, "abc", storages.Task{Content: "Do homework"})
	})

	t.Run("Scenario 12: Test add task with token from user is not exist in database", func(t *testing.T) {
		testAddTaskWithTokenFromUserIsNotExistInDb(t, storages.Task{
			Content: "userNotExists",
			UserID:  "Do homework",
		})
	})

}

func testLoginAndGetTasksWithUserHasTaskAtCreatedDate(t *testing.T, userId string, passWord string, createdDate string) {
	token := testLogin(t, userId, passWord)
	if token == nil {
		return
	}

	getListTaskPath := fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", createdDate)
	request, _ := http.NewRequest(http.MethodGet, getListTaskPath, nil)
	request.Header.Add("Authorization", *token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Need status code in get list task is 200 but the fact is %d", response.Code)
		return
	}

	getTaskResp := model.GetTaskResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &getTaskResp)
	if err != nil {
		t.Errorf("Error when parse []byte to GetTaskResponse. Detail: %v", err)
		return
	}

	if getTaskResp.Data == nil {
		t.Errorf("Expected list task in data field but the fact is different. Detail %v", response.Body)
		return
	}
}

func testLoginAndGetTasksWithUserDontHaveTaskAtCreatedDate(t *testing.T, userId string, passWord string, createdDate string) {
	token := testLogin(t, userId, passWord)
	if token == nil {
		return
	}

	getListTaskPath := fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", createdDate)
	request, _ := http.NewRequest(http.MethodGet, getListTaskPath, nil)
	request.Header.Add("Authorization", *token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Need status code in get list task is 200 but the fact is %d", response.Code)
		return
	}

	getTaskResp := model.GetTaskResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &getTaskResp)
	if err != nil {
		t.Errorf("Error when parse []byte to GetTaskResponse. Detail: %v", err)
		return
	}

	if getTaskResp.Data != nil {
		t.Errorf("Expected list task in data is empty but the fact is %v", response.Body)
		return
	}
}

func testGetTasksWithValidTokenFromUserHasTaskAtCreatedDate(t *testing.T, userId string, createdDate string) {
	token, _ := services.CreateTokenForTest(userId, time.Now().Add(time.Minute*15).Unix())
	getListTaskPath := fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", createdDate)
	request, _ := http.NewRequest(http.MethodGet, getListTaskPath, nil)
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Need status code in get list task is 200 but the fact is %d", response.Code)
		return
	}

	getTaskResp := model.GetTaskResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &getTaskResp)
	if err != nil {
		t.Errorf("Error when parse []byte to GetTaskResponse. Detail: %v", err)
		return
	}

	if getTaskResp.Data == nil {
		t.Errorf("Expected list task in data field but the fact is different. Detail %v", response.Body)
		return
	}
}

func testGetTaskWithValidTokenFromUserDontHaveTasksAtCreatedDate(t *testing.T, userId string, createdDate string) {
	token, _ := services.CreateTokenForTest(userId, time.Now().Add(time.Minute*15).Unix())
	getListTaskPath := fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", createdDate)
	request, _ := http.NewRequest(http.MethodGet, getListTaskPath, nil)
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Need status code in get list task is 200 but the fact is %d", response.Code)
		return
	}

	getTaskResp := model.GetTaskResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &getTaskResp)
	if err != nil {
		t.Errorf("Error when parse []byte to GetTaskResponse. Detail: %v", err)
		return
	}

	if getTaskResp.Data != nil {
		t.Errorf("Expected list task in data is empty but the fact is %v", response.Body)
		return
	}
}

func testGetTasksWithValidTokenButExpireFromUserHasTask(t *testing.T, userId string, createdDate string) {
	token, _ := services.CreateTokenForTest(userId, time.Now().Add(-time.Minute*15).Unix())
	getListTaskPath := fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", createdDate)
	request, _ := http.NewRequest(http.MethodGet, getListTaskPath, nil)
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 401 {
		t.Errorf("Need status code in get list task is 401 but the fact is %d", response.Code)
		return
	}
}

func testGetTaskWithValidTokenFromUserIsNotExistsInDatabase(t *testing.T, userId string, createdDate string) {
	token, _ := services.CreateTokenForTest(userId, time.Now().Add(time.Minute*15).Unix())
	getListTaskPath := fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", createdDate)
	request, _ := http.NewRequest(http.MethodGet, getListTaskPath, nil)
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Need status code in get list task is 200 but the fact is %d", response.Code)
		return
	}

	getTaskResp := model.GetTaskResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &getTaskResp)
	if err != nil {
		t.Errorf("Error when parse []byte to GetTaskResponse. Detail: %v", err)
		return
	}

	if getTaskResp.Data != nil {
		t.Errorf("Expected error but the fact is %v", response.Body)
		return
	}
}

func testLoginAndAddTask(t *testing.T, userId string, passWord string, taskInput storages.Task) {
	token := testLogin(t, userId, passWord)
	if token == nil {
		return
	}
	testAddTask(t, *token, userId, taskInput)
}

func testLoginAndAddTaskWithUserHasReachedToLimitOfTaskInDay(t *testing.T, userId string, password string, taskInput []storages.Task) {
	token := testLogin(t, userId, password)
	if token == nil {
		return
	}

	for i, _ := range taskInput {
		if i < 5 {
			passed := testAddTask(t, *token, userId, taskInput[i])
			if !passed {
				return
			}
		} else {
			passed := testAddTaskForUserWhoReachToLimit(t, *token, taskInput[i])
			if !passed {
				return
			}
		}
	}
}

/*Function support for test login and add task continuously*/
func testLogin(t *testing.T, userId string, passWord string) *string {
	loginUrl := fmt.Sprintf(config.PathConfig.PathLogin+"?user_id=%s&password=%s", userId, passWord)
	request, _ := http.NewRequest(http.MethodGet, loginUrl, nil)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Need status code in login is 200 but the fact is %d", response.Code)
		return nil
	}

	loginResp := model.LoginSuccessResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &loginResp)
	if err != nil {
		t.Errorf("Error when parse []byte to LoginSuccessResponse")
		return nil
	}

	if loginResp.Data == nil || len(*loginResp.Data) == 0 {
		t.Errorf("Need token in field data but the fact is nil")
		return nil
	}

	return loginResp.Data

}
func testAddTask(t *testing.T, token string, userId string, taskInput storages.Task) bool {
	request, _ := http.NewRequest(http.MethodPost, config.PathConfig.PathTasks, taskInput.ToIOReader())
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status code from add task action is 200 but the fact is %d", response.Code)
		return false
	}

	errorResp := model.ErrorResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &errorResp)
	if err != nil {
		t.Errorf("Error when convert []byte to ErrorResponse. Detail %v", err)
		return false
	}

	if errorResp.Error != nil {
		t.Errorf("Expected created task successfully but the fact is error. Detail: %v", errorResp)
		return false
	}

	addTaskResp := model.AddTaskResponse{}
	err = json.Unmarshal(response.Body.Bytes(), &addTaskResp)
	if err != nil {
		t.Errorf("Error when convert []byte to AddTaskResponse")
		return false
	}

	data := addTaskResp.Data
	if data == nil {
		t.Errorf("Expected data from response but the fact is nil")
		return false
	}

	if taskInput.Content != data.Content {
		t.Errorf("Field content in input and field content in task has created  is different. Field content in input is %s and field content in task has created is %s", taskInput.Content, data.Content)
		return false
	}

	// Check task is exists in list task
	request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf(config.PathConfig.PathTasks+"?created_date=%s", time.Now().Format("2006-01-02")), nil)
	request.Header.Add("Authorization", token)
	response = httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Expected status when get list tasks after add task is 200, but the fact is %d", response.Code)
		return false
	}

	getTaskResp := model.GetTaskResponse{}
	err = json.Unmarshal(response.Body.Bytes(), &getTaskResp)
	if err != nil {
		t.Errorf("Error when parse []byte to GetTaskResponse. Detail: %v", err)
		return false
	}

	if getTaskResp.Data == nil {
		t.Errorf("Expected list task in data field after add task but the fact is different. Detail %v", response.Body)
		return false
	}

	existTask := false

	for i, _ := range getTaskResp.Data {
		if getTaskResp.Data[i].Content == taskInput.Content && getTaskResp.Data[i].UserID == userId && getTaskResp.Data[i].ID == addTaskResp.Data.ID {
			existTask = true
			break
		}
	}

	if !existTask {
		t.Errorf("Not found any task we need to add in list task after add task")
		return false
	}

	return true

}
func testAddTaskForUserWhoReachToLimit(t *testing.T, token string, taskInput storages.Task) bool {
	request, _ := http.NewRequest(http.MethodPost, config.PathConfig.PathTasks, taskInput.ToIOReader())
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)

	if response.Code != 400 {
		t.Errorf("Expected status code from add task action in case user reach to limit is 400 but the fact is %d", response.Code)
		return false
	}

	errorResp := model.ErrorResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &errorResp)
	if err != nil {
		t.Errorf("Error when convert []byte to ErrorResponse. Detail %v", err)
		return false
	}

	if errorResp.Error == nil {
		t.Errorf("Expected error but the fact is nil")
		return false
	}

	return true
}

/*--------------------------------------------------------*/

func testAddTaskWithValidTokenButExpire(t *testing.T, taskInput storages.Task) {
	token, _ := services.CreateTokenForTest(taskInput.UserID, time.Now().Add(-time.Minute*15).Unix())
	request, _ := http.NewRequest(http.MethodPost, config.PathConfig.PathTasks, taskInput.ToIOReader())
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)
	if response.Code != 401 {
		t.Errorf("Need status code in get list task is 401 but the fact is %d", response.Code)
		return
	}
}

func testAddTaskWithInvalidToken(t *testing.T, token string, taskInput storages.Task) {
	request, _ := http.NewRequest(http.MethodPost, config.PathConfig.PathTasks, taskInput.ToIOReader())
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)
	if response.Code != 401 {
		t.Errorf("Need status code in get list task is 401 but the fact is %d", response.Code)
		return
	}
}

func testAddTaskWithTokenFromUserIsNotExistInDb(t *testing.T, taskInput storages.Task) {
	token, _ := services.CreateTokenForTest(taskInput.UserID, time.Now().Add(time.Minute*15).Unix())
	request, _ := http.NewRequest(http.MethodPost, config.PathConfig.PathTasks, taskInput.ToIOReader())
	request.Header.Add("Authorization", token)
	response := httptest.NewRecorder()
	services.ServiceMockForTest.ServeHTTP(response, request)
	if response.Code != 404 {
		t.Errorf("Need status code in get list task is 404 but the fact is %d", response.Code)
		return
	}

	respAddTaskErr := model.ErrorResponse{}
	err := json.Unmarshal(response.Body.Bytes(), &respAddTaskErr)
	if err != nil {
		t.Errorf("Error when convert []byte to ErrorResponse. Detail %v", err)
		return
	}

	if respAddTaskErr.Error == nil {
		t.Errorf("Expected field error in add task resp but the fact is different")
		return
	}

}
