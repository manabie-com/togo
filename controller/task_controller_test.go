package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"time"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/qgdomingo/todo-app/model"
	"github.com/qgdomingo/todo-app/mock"
	"github.com/stretchr/testify/assert"
)

func createMockTaskData(length int) []model.Task {
	taskList := []model.Task{}
	if length > 0 {
		for i := 1; i <= length; i++ {
			task := model.Task {
				ID: i,
				Title: "Sample Task Title",
				Description: "Sample Task Description",
				Username: "todo_test_user",
				CreateDate: time.Now() }
			taskList = append(taskList, task)
		}
	}
	return taskList
}

func createMockErrorMessage(message string, errMsg string) map[string]string {
	errMessage := make(map[string]string)
	errMessage["message"] = message
	errMessage["error"] = errMsg
	return errMessage
}

func createMockTaskJSON(c *gin.Context, title string, desc string, username string) ([]byte, error) {
	c.Request.Method = "POST"
    c.Request.Header.Set("Content-Type", "application/json")

	taskDetails := model.TaskUserEnteredDetails {
		Title: title,
		Description: desc,
		Username: username,
	}
	jsonbytes, err := json.Marshal(taskDetails)

	return jsonbytes, err
}

func TestFetchAllTaskAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header),
    }
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(2),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }

	taskCont.GetTasks(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchAllTaskEmptyAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header),
    }
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(0),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTasks(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchAllTaskErrorAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header),
    }
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: nil,
		ErrorMessage: createMockErrorMessage("Test Message", "Test Error Message") }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTasks(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestFetchTaskByIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "3",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(1),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchTaskByIDNotFoundAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(0),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchTaskByInvalidIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "test",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: nil,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchTaskByEmptyIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(2),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchTaskByIDErrorAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: nil,
		ErrorMessage: createMockErrorMessage("Test Message", "Test Error Message") }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestFetchTaskByUsernameAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "todo_test_user",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(2),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchTaskByUsernameNotFoundAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "todo_test_user",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: createMockTaskData(0),
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchTaskByEmptyUsernameAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: nil,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchTaskByUsernameErrorAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "todo_test_user",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: nil,
		ErrorMessage: createMockErrorMessage("Test Message", "Test Error Message") }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestCreateTaskAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: true,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestCreateTaskEmptyDetailsAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	jsonBytes, err := createMockTaskJSON(ctx, "", "", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestCreateTaskFailedInsertAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestCreateTaskErrorAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: createMockErrorMessage("Test Message", "Test Error Message") }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestUpdateTaskAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: true,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestUpdateTaskInvalidIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "test",
		},
	}

	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: true,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateTaskEmptyIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "",
		},
	}

	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: true,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateTaskEmptyDetailsAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}

	jsonBytes, err := createMockTaskJSON(ctx, "", "", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: true,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateTaskFailedUpdateAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestUpdateTaskErrorAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	// Add the JSON to the Request body
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: createMockErrorMessage("Test Message", "Test Error Message") }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestDeleteTaskAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "3",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: true,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestDeleteTaskIDNotFoundAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestDeleteTaskInvalidIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "test",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestDeleteTaskEmptyIDAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		IsTaskSuccessful: false,
		ErrorMessage: nil }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestDeleteTaskErrorAPI (t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
    ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
        Header: make(http.Header) }

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	// Create the Task Repository Mock
	taskRepoMock := mock.TaskRepositoryMock {
		TaskList: nil,
		ErrorMessage: createMockErrorMessage("Test Message", "Test Error Message") }

	taskCont := TaskController{ TaskRepo : &taskRepoMock }
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}