package apiutils

import (
	"strconv"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseSuccess(t *testing.T) {
	{
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()

		rec = httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Success case
		require.Nil(t, ResponseSuccess(c, "ok"))

		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestResponseFailure(t *testing.T) {
	{
		// Success case
		tests := []struct {
			Code int
			Err  error
		}{
			{Code: 403, Err: errors.Wrap(ErrForbidden, "403")},
			{Code: 400, Err: errors.Wrap(ErrInvalidValue, "400")},
			{Code: 503, Err: errors.Wrap(ErrIncorrectData, "503")},
			{Code: 404, Err: errors.Wrap(ErrNotFound, "404")},
			{Code: 500, Err: errors.New("something")},
		}

		for _, tt := range tests {
			t.Run(strconv.Itoa(tt.Code), func(t *testing.T) {

				e := echo.New()
				req := httptest.NewRequest(echo.GET, "/", nil)
				rec := httptest.NewRecorder()

				rec = httptest.NewRecorder()
				c := e.NewContext(req, rec)

				require.Nil(t, ResponseFailure(c, tt.Err))
				assert.Equal(t, tt.Code, rec.Code)
			})
		}
	}
}
