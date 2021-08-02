package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type loginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

type loginResponse struct {
	Data string `json:"data"`
}

func sendLoginRequest(ctx context.Context, req *loginRequest) (*loginResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", "http://localhost:5050/login", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := &loginResponse{}
	if err := json.Unmarshal(b, res); err != nil {
		return nil, err
	}
	return res, nil
}

func Test_Login_FailWithWrongUserID(t *testing.T) {
	ctx := context.Background()
	u := &loginRequest{
		UserID:   "wrong_user_id",
		Password: "wrong_password",
	}

	res, err := sendLoginRequest(ctx, u)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func Test_Login_Success(t *testing.T) {
	ctx := context.Background()
	u := &loginRequest{
		UserID:   "firstUser",
		Password: "example",
	}

	res, err := sendLoginRequest(ctx, u)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Data)
}
