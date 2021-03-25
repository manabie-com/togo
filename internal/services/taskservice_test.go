package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/configurations"
	"github.com/manabie-com/togo/internal/middleware"
	"github.com/manabie-com/togo/internal/storages"
	mockdb "github.com/manabie-com/togo/internal/storages/mock"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
)

var retrieveTasksUTCases = []postgres.RetrieveTasksParams{
	{UserID: sql.NullString{String: "firstUser", Valid: true}, CreatedDate: sql.NullString{String: "2021-03-23", Valid: true}},
	{UserID: sql.NullString{String: "secondUser", Valid: true}, CreatedDate: sql.NullString{String: "2021-03-23", Valid: true}},
}

type rtTaskResult struct {
	Success  bool
	TaskList []storages.Task
	Err      error
	Code     int
}

var rtTaskTestCases = []struct {
	RtRequest    *http.Request
	RtTaskParams postgres.RetrieveTasksParams
	RtTaskResult rtTaskResult
}{
	{
		RtRequest: &http.Request{
			Method: http.MethodGet,
			Form: map[string][]string{
				"created_date": {"2021-03-23"},
			},
		},
		RtTaskParams: postgres.RetrieveTasksParams{
			UserID:      sql.NullString{String: "firstUser", Valid: true},
			CreatedDate: sql.NullString{String: "2021-03-23", Valid: true},
		},
		RtTaskResult: rtTaskResult{
			Success: true,
			TaskList: []storages.Task{
				{
					ID:          uuid.New().String(),
					Content:     "content",
					UserID:      "firstUser",
					CreatedDate: "2021-03-23",
				},
			},
			Err:  nil,
			Code: 200,
		},
	},
	{
		RtRequest: &http.Request{
			Method: http.MethodGet,
			Form: map[string][]string{
				"created_date": {"2021-03-23"},
			},
		},
		RtTaskParams: postgres.RetrieveTasksParams{
			UserID:      sql.NullString{String: "secondUser", Valid: true},
			CreatedDate: sql.NullString{String: "2021-03-23", Valid: true},
		},
		RtTaskResult: rtTaskResult{
			Success: false,
			TaskList: []storages.Task{
				{
					ID:          uuid.New().String(),
					Content:     "content",
					UserID:      "firstUser",
					CreatedDate: "2021-03-23",
				},
			},
			Err:  sql.ErrNoRows,
			Code: 500,
		},
	},
}

func TestTaskListHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockStore := mockdb.NewMockStore(mockCtrl)
	config, err := configurations.LoadConfig("../../resources")
	require.Nil(t, err)

	for i, tc := range rtTaskTestCases {
		fmt.Println("Running TestTaskListHandler TC: ", i, ": ")
		rtTaskParams := tc.RtTaskParams
		result := tc.RtTaskResult

		mockStore.EXPECT().RetrieveTasks(gomock.Any(), rtTaskParams).Return(result.TaskList, result.Err).Times(1)
		sc := mockServiceController(t, config, mockStore)
		resp := httptest.NewRecorder()

		require.NotNil(t, sc)

		tc.RtRequest = tc.RtRequest.WithContext(context.WithValue(tc.RtRequest.Context(),
			middleware.UserAuthKey(0), tc.RtTaskParams.UserID.String))
		sc.taskListHandler(resp, tc.RtRequest)
		require.Equal(t, tc.RtTaskResult.Code, resp.Code)
	}

}

func TestUpdateTaskHandlerSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockStore := mockdb.NewMockStore(mockCtrl)
	config, err := configurations.LoadConfig("../../resources")
	require.Nil(t, err)

	mockStore.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	mockStore.EXPECT().CountTaskPerDay(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	sc := mockServiceController(t, config, mockStore)
	resp := httptest.NewRecorder()

	require.NotNil(t, sc)

	body, _ := json.Marshal(map[string]string{
		"content": "content",
	})

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	require.NoError(t, err)
	req = req.WithContext(context.WithValue(req.Context(), middleware.UserAuthKey(0), "firstUser"))
	sc.updateTasksHandler(resp, req)
	require.Equal(t, 200, resp.Code)
}

func TestUpdateTaskHandlerBadRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockStore := mockdb.NewMockStore(mockCtrl)
	config, err := configurations.LoadConfig("../../resources")
	require.Nil(t, err)

	// mockStore.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(nil).Times(1)
	// mockStore.EXPECT().CountTaskPerDay(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	sc := mockServiceController(t, config, mockStore)
	resp := httptest.NewRecorder()

	require.NotNil(t, sc)

	body, _ := json.Marshal(map[string][]string{
		"content": {"content"},
	})

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	require.NoError(t, err)
	req = req.WithContext(context.WithValue(req.Context(), middleware.UserAuthKey(0), "firstUser"))
	sc.updateTasksHandler(resp, req)
	require.Equal(t, 400, resp.Code)
}

func TestUpdateTaskHandlerInternalError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockStore := mockdb.NewMockStore(mockCtrl)
	config, err := configurations.LoadConfig("../../resources")
	require.Nil(t, err)

	mockStore.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(sql.ErrNoRows).Times(1)
	mockStore.EXPECT().CountTaskPerDay(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	sc := mockServiceController(t, config, mockStore)
	resp := httptest.NewRecorder()

	require.NotNil(t, sc)

	body, _ := json.Marshal(map[string]string{
		"content": "content",
	})

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	require.NoError(t, err)
	req = req.WithContext(context.WithValue(req.Context(), middleware.UserAuthKey(0), "firstUser"))
	sc.updateTasksHandler(resp, req)
	require.Equal(t, 500, resp.Code)
}

func TestUpdateTaskHandlerOverTask(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()

	mockStore := mockdb.NewMockStore(mockCtrl)
	config, err := configurations.LoadConfig("../../resources")
	require.Nil(t, err)

	// mockStore.EXPECT().AddTask(gomock.Any(), gomock.Any()).Return(sql.ErrNoRows).Times(1)
	mockStore.EXPECT().CountTaskPerDay(gomock.Any(), gomock.Any()).Return(int64(6), nil)
	sc := mockServiceController(t, config, mockStore)
	resp := httptest.NewRecorder()

	require.NotNil(t, sc)

	body, _ := json.Marshal(map[string]string{
		"content": "content",
	})

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewReader(body))
	require.NoError(t, err)
	req = req.WithContext(context.WithValue(req.Context(), middleware.UserAuthKey(0), "firstUser"))
	sc.updateTasksHandler(resp, req)
	require.Equal(t, 400, resp.Code)
}
