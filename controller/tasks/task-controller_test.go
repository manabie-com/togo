package tasks_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	app "github.com/manabie-com/backend/app"
	"github.com/manabie-com/backend/entity"
	mockController "github.com/manabie-com/backend/mocks/controller"
	"github.com/manabie-com/backend/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTask(t *testing.T) {
	task := entity.Task{
		ID:          uuid.NewV4().String(),
		UserID:      uuid.NewV4().String(),
		Content:     "Task1",
		Status:      "pendding",
		CreatedDate: utils.GetToday(),
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctlMock := mockController.NewMockI_TaskController(ctl)
	ctlMock.EXPECT().CreateTask(gomock.Any()).Times(1)

	server, err := app.NewServer(ctlMock)
	assert.Nil(t, nil, err)

	recorder := httptest.NewRecorder()

	data, err := json.Marshal(task)
	fmt.Println("data Exact:", task)
	if err != nil {
		fmt.Println("Error Exact:")
	}
	require.NoError(t, err)

	url := "/tasks"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)
	server.Router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
}
