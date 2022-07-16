package settings

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"manabie/todo/models"
	"manabie/todo/service/setting"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (ms *mockService) Show(ctx context.Context, memberID int) (*models.Setting, error) {
	args := ms.Called(ctx, memberID)
	return args.Get(0).(*models.Setting), args.Error(1)
}
func (ms *mockService) Create(ctx context.Context, memberID int, req *models.SettingCreateRequest) error {
	args := ms.Called(ctx, memberID, req)
	return args.Error(0)
}
func (ms *mockService) Update(ctx context.Context, settingID int, req *models.SettingUpdateRequest) error {
	args := ms.Called(ctx, settingID, req)
	return args.Error(0)
}

func Test_handler_Show(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, setting.SettingService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400",
			req: func() (*http.Request, setting.SettingService) {
				return httptest.NewRequest(echo.GET, "/users/something/settings", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, setting.SettingService) {
				mck := new(mockService)
				mck.On("Show", context.Background(), 1).Return(&models.Setting{}, errors.New("something"))

				return httptest.NewRequest(echo.GET, "/users/1/settings", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200 by mock",
			req: func() (*http.Request, setting.SettingService) {
				// Mock
				mck := new(mockService)
				mck.On("Show", context.Background(), 1).Return(&models.Setting{ID: 1}, nil)

				return httptest.NewRequest(echo.GET, "/users/1/settings", nil), mck
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
			NewSettingHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}

func Test_handler_Create(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, setting.SettingService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400 param",
			req: func() (*http.Request, setting.SettingService) {
				return httptest.NewRequest(echo.POST, "/users/something/settings", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "400 body",
			req: func() (*http.Request, setting.SettingService) {
				return httptest.NewRequest(echo.POST, "/users/1/settings", bytes.NewReader([]byte("1"))), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, setting.SettingService) {

				body := models.SettingCreateRequest{
					LimitTask: 1,
				}

				data, _ := json.Marshal(body)
				mck := new(mockService)
				mck.On("Create", context.Background(), 1, &body).Return(errors.New("something"))

				return httptest.NewRequest(echo.POST, "/users/1/settings", bytes.NewBuffer(data)), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200",
			req: func() (*http.Request, setting.SettingService) {

				body := models.SettingCreateRequest{
					LimitTask: 1,
				}

				data, _ := json.Marshal(body)
				mck := new(mockService)
				mck.On("Create", context.Background(), 1, &body).Return(nil)

				return httptest.NewRequest(echo.POST, "/users/1/settings", bytes.NewBuffer(data)), mck
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
			NewSettingHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}

func Test_handler_Update(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, setting.SettingService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "400 param",
			req: func() (*http.Request, setting.SettingService) {
				return httptest.NewRequest(echo.PUT, "/settings/something", nil), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "400 body",
			req: func() (*http.Request, setting.SettingService) {
				return httptest.NewRequest(echo.PUT, "/settings/1", bytes.NewReader([]byte("1"))), nil
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
			},
		},
		{
			name: "500",
			req: func() (*http.Request, setting.SettingService) {

				body := models.SettingUpdateRequest{
					LimitTask: 1,
				}

				data, _ := json.Marshal(body)
				mck := new(mockService)
				mck.On("Update", context.Background(), 1, &body).Return(errors.New("something"))

				return httptest.NewRequest(echo.PUT, "/settings/1", bytes.NewBuffer(data)), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200",
			req: func() (*http.Request, setting.SettingService) {

				body := models.SettingUpdateRequest{
					LimitTask: 1,
				}

				data, _ := json.Marshal(body)
				mck := new(mockService)
				mck.On("Update", context.Background(), 1, &body).Return(nil)

				return httptest.NewRequest(echo.PUT, "/settings/1", bytes.NewBuffer(data)), mck
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
			NewSettingHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}
