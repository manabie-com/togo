package transport_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/services/transport"
	"github.com/manabie-com/togo/internal/services/usecase"
	"github.com/manabie-com/togo/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	handler := transport.NewTransport()

	// create the user mock interface
	ctrl := gomock.NewController(t)
	handler.Usecase = usecase.NewMockUsecaseInterface(ctrl)

	t.Run("Success", func(t *testing.T) {
		// token
		tokenstr := fmt.Sprintf("\"%s\"\n", "tokengenerated")
		// Data
		payload := models.User{
			Username: "username123",
			Password: "password123",
		}
		data, marshalError := json.Marshal(payload)
		//do usecase method
		handler.Usecase.(*usecase.MockUsecaseInterface).EXPECT().GetAuthToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("tokengenerated", nil)
		// do the request
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		resBody, bodyErr := ioutil.ReadAll(res.Body)
		tokenResult := string(resBody)
		//testing
		assert.NoError(t, marshalError)
		assert.NoError(t, bodyErr)
		assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
		assert.Equal(t, tokenResult, tokenstr, "they should be equal")
	})

	t.Run("Input Failed", func(t *testing.T) {
		// Data
		payload := "inputstring"
		data, _ := json.Marshal(payload)
		// do the request
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		//testing
		assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	})

	t.Run("Username or Password empty", func(t *testing.T) {
		// Data
		payload := models.User{
			Username: "",
			Password: "",
		}
		data, marshalError := json.Marshal(payload)
		// do the request
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		//testing
		assert.NoError(t, marshalError)
		assert.Equal(t, res.StatusCode, http.StatusUnprocessableEntity, "they should be equal")
	})

	t.Run("Get token error", func(t *testing.T) {
		// Data
		payload := models.User{
			Username: "username123",
			Password: "password123",
		}
		data, marshalError := json.Marshal(payload)
		//do usecase method
		handler.Usecase.(*usecase.MockUsecaseInterface).EXPECT().GetAuthToken(gomock.Any(), gomock.Any(), gomock.Any()).Return("", errors.New("some errors"))
		// do the request
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		handler.Login(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		//testing
		assert.NoError(t, marshalError)
		assert.Equal(t, res.StatusCode, http.StatusInternalServerError, "they should be equal")
	})
}

func TestGetListTasks(t *testing.T) {
	t.Parallel()

	handler := transport.NewTransport()
	// create the task mock interface
	ctrl := gomock.NewController(t)
	handler.Usecase = usecase.NewMockUsecaseInterface(ctrl)

	t.Run("Sucess", func(t *testing.T) {
		allTasks := make([]*models.Task, 0)
		allTasks = append(allTasks, &models.Task{
			ID:          1,
			Content:     "statement 1",
			UserID:      1,
			CreatedDate: "2021-04-26",
		})
		allTasks = append(allTasks, &models.Task{
			ID:          2,
			Content:     "statement 2",
			UserID:      2,
			CreatedDate: "2021-04-26",
		})

		handler.Usecase.(*usecase.MockUsecaseInterface).EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(allTasks, nil)

		// create router
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/tasks", handler.ListTasks).Methods("GET")

		// create request
		req := httptest.NewRequest("GET", "/tasks", nil)
		q := req.URL.Query()
		q.Add("created_date", "2021-04-26")
		req.URL.RawQuery = q.Encode()

		// set user_id to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.UserAuthKey(0), uint64(999))
		req = req.WithContext(ctx)

		// get result tasks
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		resBody, bodyErr := ioutil.ReadAll(res.Body)

		tasks := make([]*models.Task, 0)
		unMarshalErr := json.Unmarshal(resBody, &tasks)

		assert.NoError(t, bodyErr)
		assert.NoError(t, unMarshalErr)
		assert.Equal(t, res.StatusCode, http.StatusOK, "they should be equal")
		assert.Equal(t, allTasks, tasks, "they should be equal")
	})

	t.Run("Missing Create Date", func(t *testing.T) {
		// create router
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/tasks", handler.ListTasks).Methods("GET")

		// create request
		req := httptest.NewRequest("GET", "/tasks", nil)
		q := req.URL.Query()
		q.Add("created_date", "")
		req.URL.RawQuery = q.Encode()

		// set user_id to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.UserAuthKey(0), uint64(999))
		req = req.WithContext(ctx)

		// get result tasks
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusUnprocessableEntity, "they should be equal")
	})

	t.Run("Missing UserID", func(t *testing.T) {
		// create router
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/tasks", handler.ListTasks).Methods("GET")

		// create request
		req := httptest.NewRequest("GET", "/tasks", nil)
		q := req.URL.Query()
		q.Add("created_date", "2021-09-09")
		req.URL.RawQuery = q.Encode()

		// get result tasks
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusNotFound, "they should be equal")
	})

	t.Run("Wrong Create or UserID", func(t *testing.T) {
		// create the task mock interface
		ctrl := gomock.NewController(t)
		handler.Usecase = usecase.NewMockUsecaseInterface(ctrl)
		handler.Usecase.(*usecase.MockUsecaseInterface).EXPECT().RetrieveTasks(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("some errors"))

		// create router
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/tasks", handler.ListTasks).Methods("GET")

		// create request
		req := httptest.NewRequest("GET", "/tasks", nil)
		q := req.URL.Query()
		q.Add("created_date", "2021-09-09")
		req.URL.RawQuery = q.Encode()

		// set user_id to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.UserAuthKey(0), uint64(0))
		req = req.WithContext(ctx)

		// get result tasks
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, res.StatusCode, http.StatusInternalServerError, "they should be equal")
	})
}

func TestAddTask(t *testing.T) {
	t.Parallel()
	handler := transport.NewTransport()

	// create the task mock interface
	ctrl := gomock.NewController(t)
	handler.Usecase = usecase.NewMockUsecaseInterface(ctrl)

	t.Run("Sucess", func(t *testing.T) {
		// Data
		payload := &models.Task{
			Content: "Statement 1",
		}
		data, marshalError := json.Marshal(payload)
		// do usecase method
		handler.Usecase.(*usecase.MockUsecaseInterface).EXPECT().AddTask(gomock.Any(), gomock.Any(), payload).Return(
			&models.Task{
				ID:      1,
				Content: "Statement 1",
			}, nil)
		// do the request
		req := httptest.NewRequest("POST", "/tasks", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		// set user_id to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.UserAuthKey(0), uint64(999))
		req = req.WithContext(ctx)
		//call transport method
		handler.AddTask(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		resBody, bodyErr := ioutil.ReadAll(res.Body)
		//json to struct
		task := &models.Task{}
		unMarshalErr := json.Unmarshal(resBody, &task)
		//included to facilitate the check with the returned value
		payload.ID = 1
		//testing
		assert.NoError(t, marshalError)
		assert.NoError(t, bodyErr)
		assert.NoError(t, unMarshalErr)
		assert.Equal(t, res.StatusCode, http.StatusCreated, "they should be equal")
		assert.Equal(t, task, payload, "they should be equal")
	})
	t.Run("Input Failed", func(t *testing.T) {
		// Data
		payload := "Statement 1"
		data, _ := json.Marshal(payload)
		// do the request
		req := httptest.NewRequest("POST", "/tasks", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		// call transport menthod
		handler.AddTask(w, req)
		// return response
		res := w.Result()
		defer res.Body.Close()
		// check equal
		assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	})

	t.Run("Missing UserID", func(t *testing.T) {
		// Data
		payload := &models.Task{
			Content: "Statement 1",
		}
		data, _ := json.Marshal(payload)
		// do the request
		req := httptest.NewRequest("POST", "/tasks", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		//transport method
		handler.AddTask(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		//testing
		assert.Equal(t, res.StatusCode, http.StatusNotFound, "they should be equal")
	})

	t.Run("Missing Content - Validate Addtask", func(t *testing.T) {
		// Data
		payload := &models.Task{
			Content: "",
		}
		data, _ := json.Marshal(payload)
		// do the request
		req := httptest.NewRequest("POST", "/tasks", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		// set user_id to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.UserAuthKey(0), uint64(999))
		req = req.WithContext(ctx)
		// call transport menthod
		handler.AddTask(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		// check equal
		assert.Equal(t, res.StatusCode, http.StatusUnprocessableEntity, "they should be equal")
	})

	t.Run("Task added limit reached", func(t *testing.T) {
		// Data
		payload := &models.Task{
			Content: "Statement 1",
		}
		data, marshalError := json.Marshal(payload)
		//do usecase method
		handler.Usecase.(*usecase.MockUsecaseInterface).EXPECT().AddTask(gomock.Any(), gomock.Any(), payload).Return(
			&models.Task{}, errors.New("the task daily limit is reached"))
		// do the request
		req := httptest.NewRequest("POST", "/tasks", strings.NewReader(string(data)))
		w := httptest.NewRecorder()
		// set user_id to context
		ctx := req.Context()
		ctx = context.WithValue(ctx, utils.UserAuthKey(0), uint64(999))
		req = req.WithContext(ctx)
		//call transport method
		handler.AddTask(w, req)
		//return response
		res := w.Result()
		defer res.Body.Close()
		_, bodyErr := ioutil.ReadAll(res.Body)
		//testing
		assert.NoError(t, marshalError)
		assert.NoError(t, bodyErr)
		assert.Equal(t, res.StatusCode, http.StatusBadRequest, "they should be equal")
	})
}
