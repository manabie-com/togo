package services

import (
	"encoding/json"
	"fmt"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthTokenRunner(t *testing.T) {
	t.Run("Scenario 1: Login with valid username and password => login success", func(t *testing.T) {
		testGetAuthTokenWithValidUserNameAndPassword(t, "firstUser", "example")
	})

	t.Run("Scenario 2: Login with valid username and wrong password => need login fail with status code 401", func(t *testing.T) {
		testGetAuthTokenWithInvalidInfo(t, "firstUser", "123456")
	})

	t.Run("Scenario 3: Login with wrong username and valid password => need login fail with status code 401", func(t *testing.T) {
		testGetAuthTokenWithInvalidInfo(t, "firstUser2", "example")
	})

	t.Run("Scenario 4: Login with wrong username and wrong password => need login fail with status code 401", func(t *testing.T) {
		testGetAuthTokenWithInvalidInfo(t, "firstUser2", "123456")
	})

}

func testGetAuthTokenWithValidUserNameAndPassword(t *testing.T, userName string, password string) {
	mux := http.NewServeMux()
	mux.HandleFunc(config.PathConfig.PathLogin, ServiceMockForTest.getAuthToken)

	writer := httptest.NewRecorder()

	loginPath := fmt.Sprintf(config.PathConfig.PathLogin+"?user_id=%s&password=%s", userName, password)
	request, _ := http.NewRequest(http.MethodGet, loginPath, nil)

	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Need login success but the fact is fail")
		return
	}

	loginResp := model.LoginSuccessResponse{}
	err := json.Unmarshal(writer.Body.Bytes(), &loginResp)
	if err != nil {
		t.Errorf("Error when parse []byte to LoginSuccessResponse")
		return
	}

	if loginResp.Data == nil || len(*loginResp.Data) == 0 {
		t.Errorf("Need token in field data but the fact is nil")
		return
	}

}

func testGetAuthTokenWithInvalidInfo(t *testing.T, userName string, passWord string) {
	mux := http.NewServeMux()
	mux.HandleFunc(config.PathConfig.PathLogin, ServiceMockForTest.getAuthToken)

	writer := httptest.NewRecorder()

	loginPath := fmt.Sprintf(config.PathConfig.PathLogin+"?user_id=%s&password=%s", userName, passWord)
	request, _ := http.NewRequest(http.MethodGet, loginPath, nil)

	mux.ServeHTTP(writer, request)
	if writer.Code != 401 {
		t.Errorf("Need login fail with status code 401 but the fact is %d", writer.Code)
	}

}
