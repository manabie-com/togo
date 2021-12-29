package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/controllers"
	"github.com/manabie-com/togo/entities"
	"github.com/manabie-com/togo/helpers"
	services_mocks "github.com/manabie-com/togo/mocks/services"
	"github.com/stretchr/testify/assert"
)

func TestTaskController_CreateTask(t *testing.T) {
	task := entities.Task{
		ID:          1,
		UserId:      1,
		Title:       "Task 1",
		Description: "Description Task 1",
		IsCompleted: false,
		CreatedAt:   helpers.GetDateNow(),
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	taskServiceMock := services_mocks.NewMockITaskService(ctl)

	router := gin.Default()
	gin.SetMode(gin.TestMode)

	taskController := controllers.NewTaskController(taskServiceMock)

	router.POST("/tasks", taskController.CreateTask)

	t.Run("Test created successfuly", func(t *testing.T) {
		taskServiceMock.EXPECT().CreateTask(&task).Return(nil, nil)

		bodyReq, _ := json.Marshal(task)
		request, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(bodyReq))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		respBody, err := json.Marshal(helpers.BuildResponse(true, "Task created successfully", task))
		assert.NoError(t, err)

		assert.Equal(t, 201, w.Code)
		assert.Equal(t, respBody, w.Body.Bytes())
	})

	t.Run("Test creating failed with error code 400 (Invalid body json)", func(t *testing.T) {
		userErr := errors.New("an user error")

		taskServiceMock.EXPECT().CreateTask(&task).Return(nil, userErr)

		bodyReq, _ := json.Marshal(task)
		request, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(bodyReq))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		respBody, err := json.Marshal(helpers.BuildErrorResponse("An user error occurred", userErr.Error(), nil))
		assert.NoError(t, err)

		assert.Equal(t, 400, w.Code)
		assert.Equal(t, respBody, w.Body.Bytes())
	})

	t.Run("Test creating failed with error code 400 (Invalid body json)", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte{0}))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		assert.Equal(t, 400, w.Code)
	})

	t.Run("Test creating failed with error code 500", func(t *testing.T) {
		internalErr := errors.New("an internal error")

		taskServiceMock.EXPECT().CreateTask(&task).Return(internalErr, nil)

		bodyReq, _ := json.Marshal(task)
		request, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(bodyReq))
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		respBody, err := json.Marshal(helpers.BuildErrorResponse("An error occurred", internalErr.Error(), nil))
		assert.NoError(t, err)

		assert.Equal(t, 500, w.Code)
		assert.Equal(t, respBody, w.Body.Bytes())
	})
}
