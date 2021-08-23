package integration

import (
	"encoding/json"
	"github.com/manabie-com/togo/internal/iservices"
	"github.com/manabie-com/togo/internal/tools"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("Login success", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/login?user_id=1&password=example", nil)
		require.Nil(t, err)
		w := httptest.NewRecorder()
		todoApi.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		var loginRes iservices.LoginResponse
		err = json.NewDecoder(w.Body).Decode(&loginRes)
		require.Nil(t, err)
		require.NotNil(t, loginRes.Data)
	})
	t.Run("Login fail", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/login?user_id=1&password=invalid", nil)
		require.Nil(t, err)
		w := httptest.NewRecorder()
		todoApi.ServeHTTP(w, req)
		require.Equal(t, http.StatusUnauthorized, w.Code)
		var errRes tools.TodoError
		err = json.NewDecoder(w.Body).Decode(&errRes)
		require.NotNil(t, errRes.ErrorMessage)
		require.Equal(t, "incorrect user_id/pwd", errRes.ErrorMessage)
	})
}
