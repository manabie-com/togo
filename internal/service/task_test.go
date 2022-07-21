package service

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/tonghia/togo/internal/store"
)

func (suite *ServiceTestSuite) TestRecordTask_POST() {
	testCases := []struct {
		userID                 uint64
		requestBody            string
		expectedHTTPStatusCode int
		expectedBody           string
	}{
		{
			userID:                 1,
			requestBody:            "",
			expectedHTTPStatusCode: 400,
			expectedBody:           "Invalid Request\n",
		},
	}

	// Mock Database
	suite.mockQuerier.EXPECT().GetTotalTaskByUserID(gomock.Any(), gomock.Any()).Return(store.GetTotalTaskByUserIDRow{}, nil)
	suite.mockQuerier.EXPECT().InsertTask(gomock.Any(), gomock.Any()).Return(nil, nil)

	for _, tc := range testCases {
		req := httptest.NewRequest("POST", fmt.Sprintf("/user/%d/task", tc.userID), strings.NewReader(tc.requestBody))
		w := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/user/{userID}/task", suite.service.RecordTask)
		router.ServeHTTP(w, req)

		response := w.Result()
		body, err := io.ReadAll(response.Body)

		assert.Nil(suite.T(), err)
		assert.Equal(suite.T(), tc.expectedHTTPStatusCode, response.StatusCode,
			fmt.Sprintf("Expected Response Status to be: %d, Got: %d", tc.expectedHTTPStatusCode, response.StatusCode))
		assert.Equal(suite.T(), tc.expectedBody, string(body),
			fmt.Sprintf("Expected Response body to be: %s, Got: %s", tc.expectedBody, string(body)))

	}
}
