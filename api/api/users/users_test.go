package users

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"manabie/todo/models"
	"manabie/todo/service/user"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (ms *mockService) Index(ctx context.Context) ([]*models.User, error) {
	args := ms.Called(ctx)
	return args.Get(0).([]*models.User), args.Error(1)
}

func Test_handler_Index(t *testing.T) {
	tests := []struct {
		name   string
		req    func() (*http.Request, user.UserService)
		assert func(w *httptest.ResponseRecorder, r *http.Request)
	}{
		{
			name: "500",
			req: func() (*http.Request, user.UserService) {
				mck := new(mockService)
				mck.On("Index", context.Background()).Return([]*models.User{}, errors.New("something"))

				return httptest.NewRequest(echo.GET, "/users", nil), mck
			},
			assert: func(w *httptest.ResponseRecorder, r *http.Request) {
				assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
			},
		},
		{
			name: "200 by mock",
			req: func() (*http.Request, user.UserService) {
				mck := new(mockService)
				mck.On("Index", context.Background()).Return([]*models.User{}, nil)

				return httptest.NewRequest(echo.GET, "/users", nil), mck
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
			NewUserHandler(e, sv)

			e.ServeHTTP(w, r)
			tt.assert(w, r)
		})
	}
}
