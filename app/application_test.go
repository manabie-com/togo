package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/backend/entity"
	mockController "github.com/manabie-com/backend/mocks/controller"
	"github.com/manabie-com/backend/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
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

	server, err := NewServer(ctlMock)
	assert.NoError(t, err)

	assert.Nil(t, nil, err)

	recorder := httptest.NewRecorder()

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Error Exact:")
	}
	assert.NoError(t, err)

	url := "/tasks"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	assert.NoError(t, err)
	server.Router.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)

}
