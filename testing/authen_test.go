package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"togo/utils"

	"github.com/stretchr/testify/assert"
)

type AuthenResponse struct {
	Data    DataResponse `json:"data"`
	Error   interface{}  `json:"error"`
	Message string       `json:"message"`
	Success bool         `json:"success"`
}

type DataResponse struct {
	Token string `json:"token"`
}

type AuthenData struct {
	endPoint       string
	method         string
	input          string
	expectedMsg    string
	expectedStatus int
	resp           *AuthenResponse
}

func (au *AuthenData) Test(t *testing.T) {
	if au == nil {
		t.Error("initial data for this method is missing")
	}
	url := fmt.Sprintf("%s%s", path, au.endPoint)
	jsonStr := bytes.NewBufferString(au.input)
	req, err := http.NewRequest(au.method, url, jsonStr)
	if err != nil {
		t.Error("network error: ", err)
	}
	req.Header.Set("content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Got error on api request %s", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var record AuthenResponse

	unmarshalErr := json.Unmarshal(body, &record)
	if unmarshalErr != nil {
		t.Errorf("got error on response parse: %s", unmarshalErr)
	}

	au.resp = &record
	assert.Equal(t, au.expectedStatus, resp.StatusCode)
	assert.Equal(t, au.expectedMsg, au.resp.Message)
}
func LoginTest(t *testing.T) {

	t.Run("PasswordNotValid", func(t *testing.T) {
		data := &AuthenData{
			endPoint:       "/auth/login",
			method:         "POST",
			input:          `{"email":"test@test.com", "password":"123"}`,
			expectedMsg:    `Password is at least 6 characters`,
			expectedStatus: http.StatusBadRequest,
		}
		data.Test(t)
		respData := data.resp
		assert.NotNil(t, respData.Error)
	})

	t.Run("WrongUsernameOrPassword", func(t *testing.T) {
		data := &AuthenData{
			endPoint:       "/auth/login",
			method:         "POST",
			input:          `{"email":"test@test.com", "password":"12345678"}`,
			expectedMsg:    `The username or password was not correct`,
			expectedStatus: http.StatusUnauthorized,
		}
		data.Test(t)
	})

	t.Run("AuthenSuccess", func(t *testing.T) {
		data := &AuthenData{
			endPoint:       "/auth/login",
			method:         "POST",
			input:          `{"email":"admin@gmail.com", "password":"123456"}`,
			expectedMsg:    `Success`,
			expectedStatus: http.StatusOK,
		}
		data.Test(t)
		respData := data.resp
		j := utils.NewJWT()
		_, err := j.ParseToken(respData.Data.Token)
		assert.Equal(t, nil, err)
	})
}
