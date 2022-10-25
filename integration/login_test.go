package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"togo.com/pkg/model"
)

func TestLogin(t *testing.T) {
	t.Run("login ok", func(t *testing.T) {
		url := fmt.Sprintf("%s%s", serverUrl, "/login")
		reqLogin := model.LoginRequest{
			UserName: "firstUser",
			Password: "example",
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqLogin)
		assert.NoError(t, err)
		client := &http.Client{}
		req, err := http.NewRequest(echo.POST, url, &buf)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		assert.NoError(t, err)
		byteBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		responseData := struct {
			Status  int64  `json:"status"`
			Message string `json:"message"`
			Data    struct {
				Token string `json:"token"`
			} `json:"data"`
		}{}
		err = json.Unmarshal(byteBody, &responseData)
		assert.NoError(t, err)
		assert.NotNil(t, responseData.Data.Token)

		assert.EqualValues(t, http.StatusOK, res.StatusCode)
		res.Body.Close()

	})
	t.Run("login not ok", func(t *testing.T) {
		url := fmt.Sprintf("%s%s", serverUrl, "/login")
		reqLogin := model.LoginRequest{
			UserName: "user",
			Password: "pass",
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(reqLogin)
		assert.NoError(t, err)
		client := &http.Client{}
		req, err := http.NewRequest(echo.POST, url, &buf)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		assert.NoError(t, err)
		byteBody, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)
		responseData := struct {
			Status  int64  `json:"status"`
			Message string `json:"message"`
			Data    struct {
				Token string `json:"token"`
			} `json:"data"`
		}{}
		err = json.Unmarshal(byteBody, &responseData)
		assert.NoError(t, err)
		assert.Equal(t, `Error :invalid user`, responseData.Message)
		assert.EqualValues(t, http.StatusOK, res.StatusCode)
		res.Body.Close()

	})
}
