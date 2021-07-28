package transport_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/mocks"
	errutil "github.com/manabie-com/togo/internal/pkg/error_utils"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/transport"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	testCases := []struct {
		context      string
		buildStubs   func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface)
		user         storages.User
		params       url.Values
		expectedErr  errutil.MessageErr
		expectedCode int
	}{
		{
			context: "success",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
				userUsecase.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(nil)
				os.Setenv("SECRET_KEY_JWT", "secret_key")
			},
			params:       url.Values{"user_id": []string{"firstUser"}},
			expectedCode: http.StatusOK,
		},
		{
			context: "invalid user or password",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
				userUsecase.EXPECT().ValidateUser(gomock.Any(), gomock.Any()).Return(errors.New("invalid user or password"))
			},
			params:       url.Values{"user_id": []string{"firstUser"}},
			expectedCode: http.StatusUnauthorized,
			expectedErr:  errutil.NewUnauthorizedError("invalid user or password"),
		},
	}

	for _, c := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		t.Run(c.context, func(t *testing.T) {
			userUsecase := mocks.NewMockUserUsecaseInterface(controller)
			c.buildStubs(userUsecase, nil)

			transport := transport.Transport{
				UserUsecase: userUsecase,
			}

			router := mux.NewRouter()
			router.HandleFunc("/login", transport.Login).Methods("GET")
			req, _ := http.NewRequest(http.MethodGet, "/login", nil)
			req.URL.RawQuery = c.params.Encode()

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assert.Equal(t, c.expectedCode, rr.Code)
			if c.expectedCode != http.StatusOK {
				message, err := errutil.NewApiErrFromBytes(rr.Body.Bytes())
				assert.NoError(t, err)
				assert.Equal(t, c.expectedErr, message)
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	testCases := []struct {
		context      string
		buildStubs   func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface)
		params       interface{}
		expectedErr  errutil.MessageErr
		expectedCode int
		userID       interface{}
	}{
		{
			context: "success",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
				userUsecase.EXPECT().CreateTask(gomock.Any()).Return(nil)
			},
			params:       storages.Task{Content: "test task"},
			expectedCode: http.StatusOK,
			userID:       "1111",
		},
		{
			context: "parse params error",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
			},
			params:       `example`,
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("json: cannot unmarshal string into Go value of type storages.Task"),
			userID:       "1111",
		},
		{
			context: "user_id is empty",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
			},
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("user_id is empty"),
			userID:       nil,
		},
		{
			context: "create task fail",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
				userUsecase.EXPECT().CreateTask(gomock.Any()).Return(errors.New("create task fail"))
			},
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("create task fail"),
			userID:       "1111",
		},
	}

	for _, c := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		t.Run(c.context, func(t *testing.T) {
			userUsecase := mocks.NewMockUserUsecaseInterface(controller)
			c.buildStubs(userUsecase, nil)

			transport := transport.Transport{
				UserUsecase: userUsecase,
			}

			jsonStr, err := json.Marshal(c.params)
			assert.NoError(t, err)

			router := mux.NewRouter()
			router.HandleFunc("/tasks", transport.CreateTask).Methods("POST")

			req, err := http.NewRequestWithContext(context.WithValue(context.Background(), 0, c.userID), http.MethodPost, "/tasks", bytes.NewBuffer(jsonStr))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, c.expectedCode, rr.Code)

			if c.expectedCode != http.StatusOK {
				message, err := errutil.NewApiErrFromBytes(rr.Body.Bytes())
				assert.NoError(t, err)
				assert.Equal(t, c.expectedErr, message)
			}
		})
	}
}

var (
	task = storages.Task{
		Content: "task 1",
	}
)

func TestListTasks(t *testing.T) {
	testCases := []struct {
		context      string
		buildStubs   func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface)
		params       url.Values
		expectedErr  errutil.MessageErr
		expectedCode int
		userID       interface{}
	}{
		{
			context: "success",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
				taskUsecase.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any()).Return([]*storages.Task{&task}, nil)
			},
			params:       url.Values{"created_date": []string{"2021-07-27"}},
			expectedCode: http.StatusOK,
			userID:       "1111",
		},
		{
			context: "cannot get user_id",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
			},
			params:       url.Values{"created_date": []string{"2021-07-27"}},
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("user_id is empty"),
			userID:       nil,
		},
		{
			context: "retrieve tasks fail",
			buildStubs: func(userUsecase *mocks.MockUserUsecaseInterface, taskUsecase *mocks.MockTaskUsecaseInterface) {
				taskUsecase.EXPECT().RetrieveTasks(gomock.Any(), gomock.Any()).Return([]*storages.Task{&task}, errors.New("retrieve tasks fail"))
			},
			params:       url.Values{"created_date": []string{"2021-07-27"}},
			expectedCode: http.StatusBadRequest,
			expectedErr:  errutil.NewBadRequestError("retrieve tasks fail"),
			userID:       "1111",
		},
	}

	for _, c := range testCases {
		controller := gomock.NewController(t)
		defer controller.Finish()

		t.Run(c.context, func(t *testing.T) {
			taskUsecase := mocks.NewMockTaskUsecaseInterface(controller)
			c.buildStubs(nil, taskUsecase)
			transport := transport.Transport{
				TaskUsecase: taskUsecase,
			}

			router := mux.NewRouter()
			router.HandleFunc("/tasks", transport.ListTasks).Methods("GET")

			req, err := http.NewRequestWithContext(context.WithValue(context.Background(), 0, c.userID), http.MethodGet, "/tasks", nil)
			assert.NoError(t, err)
			req.URL.RawQuery = c.params.Encode()

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, c.expectedCode, rr.Code)

			if c.expectedCode != http.StatusOK {
				message, err := errutil.NewApiErrFromBytes(rr.Body.Bytes())
				assert.NoError(t, err)
				assert.Equal(t, c.expectedErr, message)
			}
		})
	}
}
