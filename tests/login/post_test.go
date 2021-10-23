package login

import (
	"fmt"
	"net/http"
	"testing"
)

func TestUserLogin(t *testing.T) {
	type testcase struct {
		name       string
		userID     string
		password   string
		statusCode int
	}

	tcs := []testcase{
		{"User should login successfully with valid userId + password",
			"firstUser",
			"example",
			http.StatusOK},
		{"User should try to login with invalid userId should fail",
			"firstID",
			"example",
			http.StatusUnauthorized},
		{"User should try to login with invalid password should fail",
			"firstID",
			"hihi",
			http.StatusUnauthorized},
	}

	runInParallel := func(t *testing.T, tc testcase) {
		t.Parallel()
		resp, err := http.Get(fmt.Sprintf("http://localhost:5050/login?user_id=%s&password=%s", tc.userID, tc.password))
		if err != nil {
			t.Fatalf("Error happen while trying to login: %v", err)
		}

		if resp.StatusCode != tc.statusCode {
			t.Fatalf("Test case got failed: %s, expected statusCode: %d", tc.name, tc.statusCode)
		}
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			runInParallel(t, tc)
		})
	}
}