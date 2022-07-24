package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/trangmaiq/togo/pkg/uuidx"
)

func TestHandler_CreateTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	t.Run("should return 400 if the given request is not json", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader("bug"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		persister := NewMockPersister(ctrl)
		persister.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Times(0)

		d := NewMockDependencies(ctrl)
		d.EXPECT().Persister().Times(0)

		h := New(d)

		err := h.CreateTasks()(c)
		require.Error(t, err)
		require.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 400 if the given user id is invalid", func(t *testing.T) {
		input := CreateTaskRequest{
			UserID: "no one",
			Title:  "introduce togo service",
			Note:   "should write better doc",
		}

		b, err := json.Marshal(&input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		persister := NewMockPersister(ctrl)
		persister.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Times(0)

		d := NewMockDependencies(ctrl)
		d.EXPECT().Persister().Times(0)

		h := New(d)

		err = h.CreateTasks()(c)
		require.ErrorIs(t, err, ErrInvalidUserID)
		require.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 400 if the given task title is empty", func(t *testing.T) {
		input := CreateTaskRequest{
			UserID: uuidx.UUID(uuid.NewString()),
			Title:  "",
			Note:   "should write better doc",
		}

		b, err := json.Marshal(&input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		persister := NewMockPersister(ctrl)
		persister.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Times(0)

		d := NewMockDependencies(ctrl)
		d.EXPECT().Persister().Times(0)

		h := New(d)

		err = h.CreateTasks()(c)
		require.ErrorIs(t, err, ErrEmptyTitle)
		require.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 429 if user makes a lot of requests in a short period of time", func(t *testing.T) {
		input := CreateTaskRequest{
			UserID: uuidx.UUID(uuid.NewString()),
			Title:  "introduce togo service",
			Note:   "should write better doc",
		}

		b, err := json.Marshal(&input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		limiter := NewMockRateLimiter(ctrl)
		limiter.EXPECT().
			AllowN(gomock.Any(), input.UserID.String(), int64(1)).
			Return(false).
			Times(1)

		d := NewMockDependencies(ctrl)
		d.EXPECT().
			RateLimiter().
			Return(limiter).
			Times(1)

		h := New(d)

		err = h.CreateTasks()(c)
		require.Equal(t, http.StatusTooManyRequests, rec.Code)

		var res RegularErrorResponse
		err = json.NewDecoder(rec.Body).Decode(&res)
		require.NoError(t, err)
		require.Equal(t, RegularErrorResponse{
			Error:            "too many requests",
			ErrorDescription: "a lot of requests are made in a short period of time",
		}, res)
	})

	t.Run("should return 500 if cannot store task to persistence storage", func(t *testing.T) {
		input := CreateTaskRequest{
			UserID: uuidx.UUID(uuid.NewString()),
			Title:  "introduce togo service",
			Note:   "should write better doc",
		}

		b, err := json.Marshal(&input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		errDB := errors.New("db error")
		persister := NewMockPersister(ctrl)
		persister.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(errDB).Times(1)

		limiter := NewMockRateLimiter(ctrl)
		limiter.EXPECT().
			AllowN(gomock.Any(), input.UserID.String(), int64(1)).
			Return(true).
			Times(1)

		d := NewMockDependencies(ctrl)
		d.EXPECT().
			RateLimiter().
			Return(limiter).
			Times(1)
		d.EXPECT().
			Persister().
			Return(persister).
			Times(1)

		h := New(d)

		err = h.CreateTasks()(c)
		require.ErrorIs(t, err, errDB)
		require.Equal(t, http.StatusInternalServerError, rec.Code)

		var res RegularErrorResponse
		err = json.NewDecoder(rec.Body).Decode(&res)
		require.NoError(t, err)
		require.Equal(t, RegularErrorResponse{
			Error:            "create task failed",
			ErrorDescription: "cannot store task to persistence storage",
		}, res)
	})

	t.Run("should return 200", func(t *testing.T) {
		input := CreateTaskRequest{
			UserID: uuidx.UUID(uuid.NewString()),
			Title:  "introduce togo service",
			Note:   "should write better doc",
		}

		b, err := json.Marshal(&input)
		require.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		persister := NewMockPersister(ctrl)
		persister.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		limiter := NewMockRateLimiter(ctrl)
		limiter.EXPECT().
			AllowN(gomock.Any(), input.UserID.String(), int64(1)).
			Return(true).
			Times(1)

		d := NewMockDependencies(ctrl)
		d.EXPECT().
			RateLimiter().
			Return(limiter).
			Times(1)
		d.EXPECT().Persister().Return(persister).Times(1)

		h := New(d)

		err = h.CreateTasks()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)

		var res CreateTaskResponse
		err = json.NewDecoder(rec.Body).Decode(&res)
		require.NoError(t, err)
		require.NotEmpty(t, res.ID)
	})
}
