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
		userTotalTask          int64
		userDailyLimit         uint32
		requestBody            string
		expectedHTTPStatusCode int
		expectedBody           string
	}{

		{
			userID:                 0,
			requestBody:            "",
			expectedHTTPStatusCode: 400,
			expectedBody:           "Invalid user_id\n",
		},
		{
			userID:                 1,
			requestBody:            "",
			expectedHTTPStatusCode: 400,
			expectedBody:           "Invalid Request\n",
		},
		{
			userID:                 2,
			userTotalTask:          3,
			userDailyLimit:         1,
			requestBody:            "{\"name\": \"read book\"}",
			expectedHTTPStatusCode: 403,
			expectedBody:           "Forbidden: maximum daily task reached\n",
		},
		{
			userID:                 3,
			userTotalTask:          1,
			userDailyLimit:         3,
			requestBody:            "{\"name\": \"read book\"}",
			expectedHTTPStatusCode: 200,
			expectedBody:           "{\"message\":\"Success\"}",
		},
	}

	// Mock Database
	suite.mockQuerier.EXPECT().InsertTask(gomock.Any(), gomock.Any()).Return(nil, nil)

	for _, tc := range testCases {
		// Mock db GetTotalTaskByUserID
		suite.mockQuerier.EXPECT().GetTotalTaskByUserID(gomock.Any(), tc.userID).Return(store.GetTotalTaskByUserIDRow{TotalTask: tc.userTotalTask}, nil)

		// Mock user limit svc GetUserLimit
		// fmt.Println("tc", tc.userDailyLimit)
		suite.mockUserLimitSvc.EXPECT().GetUserLimit(tc.userID).Return(tc.userDailyLimit)

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

func (suite *ServiceTestSuite) TestRecordTask_MethodNotAllow() {
	testCases := []struct {
		userID                 uint64
		requestBody            string
		expectedHTTPStatusCode int
		expectedBody           string
	}{
		{
			userID:                 1,
			requestBody:            "",
			expectedHTTPStatusCode: 405,
			expectedBody:           "Method not allowed\n",
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest("PUT", fmt.Sprintf("/user/%d/task", tc.userID), strings.NewReader(tc.requestBody))
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
