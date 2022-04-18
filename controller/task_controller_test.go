package controller

import (
	"testing"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Unit Testing for the Task Controller

func TestFetchAllTaskAPI (t *testing.T) {
	w, ctx := createRegularContext()
	
	taskCont := getTaskController(createMockTaskData(2), false, nil)

	taskCont.GetTasks(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchAllTaskEmptyAPI (t *testing.T) {
	w, ctx := createRegularContext()
	
	taskCont := getTaskController(createMockTaskData(0), false, nil)
	
	taskCont.GetTasks(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchAllTaskErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	taskCont := getTaskController(nil, false, createMockErrorMessage("Test Message", "Test Error Message"))
	
	taskCont.GetTasks(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestFetchTaskByIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "3",
		},
	}

	taskCont := getTaskController(createMockTaskData(1), false, nil)
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchTaskByIDNotFoundAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	taskCont := getTaskController(createMockTaskData(0), false, nil)
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchTaskByInvalidIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "test",
		},
	}
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchTaskByEmptyIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "",
		},
	}
	
	taskCont := getTaskController(createMockTaskData(2), false, nil)
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchTaskByIDErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	taskCont := getTaskController(nil, false, createMockErrorMessage("Test Message", "Test Error Message"))
	
	taskCont.GetTaskById(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestFetchTaskByUsernameAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "todo_test_user",
		},
	}
	
	taskCont := getTaskController(createMockTaskData(2), false, nil)
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchTaskByUsernameNotFoundAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "todo_test_user",
		},
	}
	
	taskCont := getTaskController(createMockTaskData(0), false, nil)
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchTaskByEmptyUsernameAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "",
		},
	}
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchTaskByUsernameErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "user",
			Value: "todo_test_user",
		},
	}
	
	taskCont := getTaskController(nil, false, createMockErrorMessage("Test Message", "Test Error Message"))
	
	taskCont.GetTaskByUser(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestCreateTaskAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, true, nil)
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestCreateTaskEmptyDetailsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockTaskJSON(ctx, "", "", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestCreateTaskFailedInsertAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestCreateTaskErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockTaskJSON(ctx, "Test Task Title", "Test Task Description", "todo_task_user")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, false, createMockErrorMessage("Test Message", "Test Error Message"))
	
	taskCont.CreateTask(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestUpdateTaskAPI (t *testing.T) {
	w, ctx := createRegularContext()

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

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, true, nil)
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestUpdateTaskInvalidIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "test",
		},
	}

	taskCont := getTaskController(nil, true, nil)
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateTaskEmptyIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "",
		},
	}

	taskCont := getTaskController(nil, true, nil)
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateTaskEmptyDetailsAPI (t *testing.T) {
	w, ctx := createRegularContext()

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

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, true, nil)
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateTaskFailedUpdateAPI (t *testing.T) {
	w, ctx := createRegularContext()

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

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code OK")
}

func TestUpdateTaskErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

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

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))
	
	taskCont := getTaskController(nil, false, createMockErrorMessage("Test Message", "Test Error Message"))
	
	taskCont.UpdateTask(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestDeleteTaskAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "3",
		},
	}

	taskCont := getTaskController(nil, true, nil)
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestDeleteTaskIDNotFoundAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestDeleteTaskInvalidIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "test",
		},
	}
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestDeleteTaskEmptyIDAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "",
		},
	}
	
	taskCont := getTaskController(nil, false, nil)
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestDeleteTaskErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "id",
			Value: "1",
		},
	}
	
	taskCont := getTaskController(nil, false, createMockErrorMessage("Test Message", "Test Error Message"))
	
	taskCont.DeleteTask(ctx)

	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}