//go:build integration
// +build integration

package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RegisterOrLogin(t *testing.T) {
	body := `{"username":"admin","password":"admin@@"}`
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/v1/users", strings.NewReader(body))
	assert.Nil(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	rawResp, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var jsonData map[string]interface{}
	err = json.Unmarshal(rawResp, &jsonData)
	assert.Nil(t, err)

	token := jsonData["token"]
	assert.NotEmpty(t, token)
}

func Test_CreateTask(t *testing.T) {
	// get token
	body := `{"username":"admin","password":"admin@@"}`
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/v1/users", strings.NewReader(body))
	assert.Nil(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	rawResp, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	var jsonData map[string]interface{}
	err = json.Unmarshal(rawResp, &jsonData)
	assert.Nil(t, err)

	token := jsonData["token"].(string)
	assert.NotEmpty(t, token)

	// create task
	t.Run("create tast", func(t *testing.T) {
		body := `{"title":"the title","content":"anything"}`
		req, err := http.NewRequest(http.MethodPost, "http://localhost:9090/v1/tasks", strings.NewReader(body))
		req.Header.Add("authorization", "Bearer "+token)
		assert.Nil(t, err)

		resp, err := http.DefaultClient.Do(req)
		assert.Nil(t, err)
		defer resp.Body.Close()

		rawResp, err := io.ReadAll(resp.Body)
		assert.Nil(t, err)

		var jsonData map[string]interface{}
		err = json.Unmarshal(rawResp, &jsonData)
		assert.Nil(t, err)
		fmt.Println(jsonData)
	})

}
