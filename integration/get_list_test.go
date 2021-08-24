package integration

import (
	"encoding/json"
	"github.com/manabie-com/togo/internal/iservices"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetListTask(t *testing.T) {
	t.Run("Get list success with data", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/login?user_id=1&password=example", nil)
		require.Nil(t, err)
		w := httptest.NewRecorder()
		todoApi.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		var loginRes iservices.LoginResponse
		err = json.NewDecoder(w.Body).Decode(&loginRes)
		require.Nil(t, err)
		require.NotNil(t, loginRes.Data)

		reqGetList, err := http.NewRequest(http.MethodGet, "/tasks?created_date=2021-08-20", nil)
		require.Nil(t, err)
		reqGetList.Header.Set("Authorization", loginRes.Data)
		wGetList := httptest.NewRecorder()
		todoApi.ServeHTTP(wGetList, reqGetList)
		require.Equal(t, http.StatusOK, wGetList.Code)
		var listTaskRes iservices.ListTaskResponse
		err = json.NewDecoder(wGetList.Body).Decode(&listTaskRes)
		require.Nil(t, err)
		require.NotNil(t, listTaskRes.Data)
		require.Equal(t, 4, len(listTaskRes.Data))
	})

	t.Run("Get list success with data nil", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/login?user_id=1&password=example", nil)
		require.Nil(t, err)
		w := httptest.NewRecorder()
		todoApi.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		var loginRes iservices.LoginResponse
		err = json.NewDecoder(w.Body).Decode(&loginRes)
		require.Nil(t, err)
		require.NotNil(t, loginRes.Data)

		reqGetList, err := http.NewRequest(http.MethodGet, "/tasks?created_date=2021-08-23", nil)
		require.Nil(t, err)
		reqGetList.Header.Set("Authorization", loginRes.Data)
		wGetList := httptest.NewRecorder()
		todoApi.ServeHTTP(wGetList, reqGetList)
		require.Equal(t, http.StatusOK, wGetList.Code)
		var listTaskRes iservices.ListTaskResponse
		err = json.NewDecoder(wGetList.Body).Decode(&listTaskRes)
		require.Nil(t, err)
		require.Nil(t, listTaskRes.Data)
	})

	t.Run("Get list with invalid token", func(t *testing.T) {
		reqGetList, err := http.NewRequest(http.MethodGet, "/tasks?created_date=2021-08-23", nil)
		require.Nil(t, err)
		reqGetList.Header.Set("Authorization", "12112121212")
		wGetList := httptest.NewRecorder()
		todoApi.ServeHTTP(wGetList, reqGetList)
		require.Equal(t, http.StatusInternalServerError, wGetList.Code)
	})
}
