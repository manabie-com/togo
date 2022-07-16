package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"manabie/todo/models"
	"manabie/todo/service/task"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (ms *mockService) Index(ctx context.Context, memberId int, date string) ([]*models.Task, error) {
	args := ms.Called(ctx, memberId, date)
	return args.Get(0).([]*models.Task), args.Error(1)
}
func (ms *mockService) Show(ctx context.Context, ID int) (*models.Task, error) {
	args := ms.Called(ctx, ID)
	return args.Get(0).(*models.Task), args.Error(1)
}
func (ms *mockService) Create(ctx context.Context, memberID int, tk *models.TaskCreateRequest) error {
	args := ms.Called(ctx, memberID, tk)
	return args.Error(0)
}
func (ms *mockService) Update(ctx context.Context, ID int, tk *models.Task) error {
	args := ms.Called(ctx, ID, tk)
	return args.Error(0)
}
func (ms *mockService) Delete(ctx context.Context, taskID int) error {
	args := ms.Called(ctx, taskID)
	return args.Error(0)
}

func Test_handler_Index(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, task.TaskService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.GET, "/users/something/tasks", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, task.TaskService) {
				mck := new(mockService)
				mck.On("Index", context.Background(), 1, "").Return([]*models.Task{}, errors.New("something"))

				return httptest.NewRequest(echo.GET, "/users/1/tasks", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200 by mock",
			req: func() (*http.Request, task.TaskService) {
				mck := new(mockService)
				mck.On("Index", context.Background(), 1, "").Return([]*models.Task{}, nil)

				return httptest.NewRequest(echo.GET, "/users/1/tasks", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusOK, w.Result().StatusCode)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := echo.New()
			r, sv := tt.req()
			w := httptest.NewRecorder()

			// New Handler
			NewTaskHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}

func Test_handler_Show(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, task.TaskService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.GET, "/tasks/something", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, task.TaskService) {
				mck := new(mockService)
				mck.On("Show", context.Background(), 1).Return(&models.Task{}, errors.New("something"))

				return httptest.NewRequest(echo.GET, "/tasks/1", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200",
			req: func() (*http.Request, task.TaskService) {
				mck := new(mockService)
				mck.On("Show", context.Background(), 1).Return(&models.Task{}, nil)

				return httptest.NewRequest(echo.GET, "/tasks/1", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusOK, w.Result().StatusCode)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := echo.New()
			r, sv := tt.req()
			w := httptest.NewRecorder()

			// New Handler
			NewTaskHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}

func Test_handler_Create(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, task.TaskService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400 param",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.POST, "/users/something/tasks", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "400 body",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.POST, "/users/1/tasks", bytes.NewReader([]byte("1"))), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, task.TaskService) {
				body := models.TaskCreateRequest{
					Content:    "something",
					TargetDate: "2022-07-15",
				}

				data, _ := json.Marshal(body)

				mck := new(mockService)
				mck.On("Create", context.Background(), 1, &body).Return(errors.New("something"))

				return httptest.NewRequest(echo.POST, "/users/1/tasks", bytes.NewReader(data)), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200",
			req: func() (*http.Request, task.TaskService) {
				body := models.TaskCreateRequest{
					Content:    "something",
					TargetDate: "2022-07-15",
				}

				data, _ := json.Marshal(body)

				mck := new(mockService)
				mck.On("Create", context.Background(), 1, &body).Return(nil)

				return httptest.NewRequest(echo.POST, "/users/1/tasks", bytes.NewReader(data)), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusOK, w.Result().StatusCode)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := echo.New()
			r, sv := tt.req()
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			// New Handler
			NewTaskHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}

func Test_handler_Update(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, task.TaskService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400 param",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.PUT, "/tasks/something", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "400 body",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.PUT, "/tasks/1", bytes.NewReader([]byte("1"))), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, task.TaskService) {
				body := models.Task{}

				data, _ := json.Marshal(body)

				mck := new(mockService)
				mck.On("Update", context.Background(), 1, &body).Return(errors.New("something"))

				return httptest.NewRequest(echo.PUT, "/tasks/1", bytes.NewReader(data)), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200",
			req: func() (*http.Request, task.TaskService) {
				body := models.Task{
					Content: "something",
				}

				data, _ := json.Marshal(body)

				mck := new(mockService)
				mck.On("Update", context.Background(), 1, &body).Return(nil)

				return httptest.NewRequest(echo.PUT, "/tasks/1", bytes.NewReader(data)), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusOK, w.Result().StatusCode)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := echo.New()
			r, sv := tt.req()
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			// New Handler
			NewTaskHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}

func Test_handler_Delete(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, task.TaskService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400 param",
			req: func() (*http.Request, task.TaskService) {
				return httptest.NewRequest(echo.DELETE, "/tasks/something", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, task.TaskService) {
				mck := new(mockService)
				mck.On("Delete", context.Background(), 1).Return(errors.New("something"))

				return httptest.NewRequest(echo.DELETE, "/tasks/1", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200",
			req: func() (*http.Request, task.TaskService) {

				mck := new(mockService)
				mck.On("Delete", context.Background(), 1).Return(nil)

				return httptest.NewRequest(echo.DELETE, "/tasks/1", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusOK, w.Result().StatusCode)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			e := echo.New()
			r, sv := tt.req()
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			// New Handler
			NewTaskHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}
