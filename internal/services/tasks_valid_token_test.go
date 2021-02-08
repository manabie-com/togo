package services

import (
	"github.com/manabie-com/togo/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var ValidToken, _ = CreateTokenForTest("firstUser", time.Now().Add(time.Minute*15).Unix())
var ValidTokenButExpire, _ = CreateTokenForTest("firstUser", time.Now().Add(-time.Minute*15).Unix())
var InvalidToken = "abcdefgh123456"

func TestValidTokenRunner(t *testing.T) {
	t.Run("Scenario 1: Test with valid token => Expect return true", func(t *testing.T) {
		testWithValidToken(t, ValidToken)
	})

	t.Run("Scenario 2: Test with valid token but expire => Expect return false", func(t *testing.T) {
		testWithValidTokenButExpire(t, ValidTokenButExpire)
	})

	t.Run("Scenario 3: Test with invalid token => Expect return false", func(t *testing.T) {
		testWithInvalidToken(t, InvalidToken)
	})

}

func testWithValidToken(t *testing.T, token string) {
	request := httptest.NewRequest(http.MethodGet, config.PathConfig.PathTasks, nil)
	request.Header.Set("Authorization", token)
	_, ok := ServiceMockForTest.validToken(request)
	if !ok {
		t.Error("Expect value is true, but the fact is false")
	}
}

func testWithValidTokenButExpire(t *testing.T, token string) {
	request := httptest.NewRequest(http.MethodGet, config.PathConfig.PathTasks, nil)
	request.Header.Set("Authorization", token)
	_, ok := ServiceMockForTest.validToken(request)
	if ok {
		t.Error("Expect value is false, but the fact is true")
	}
}

func testWithInvalidToken(t *testing.T, token string) {
	request := httptest.NewRequest(http.MethodGet, config.PathConfig.PathTasks, nil)
	request.Header.Set("Authorization", token)
	_, ok := ServiceMockForTest.validToken(request)
	if ok {
		t.Error("Expect value is false, but the fact is true")
	}
}
