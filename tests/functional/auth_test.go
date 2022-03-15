package functional_test

import (
	"strings"
	"testing"
	_testUtil "togo-service/tests/utils"

	"github.com/gofiber/fiber/v2/utils"
)

const REGISTER_URI = "/api/v1/register"
const LOGIN_URI = "/api/v1/login"

func Test_Auth(t *testing.T) {
	t.Run("[POST] Register - Successful", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"username": "abc",
			"password": "123"
		}`)
		req := _testUtil.ApiRoute("POST", REGISTER_URI, bodyReader, "")
		resp, err := app.Test(req, -1)
		if err != nil {
			println(err.Error())
		}
		bodyR, _ := _testUtil.GetBody(resp)
		expected := `{"error":false,"message":"Created user successfully"}`
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 200 Result")
	})

	t.Run("[POST] Register - Duplicated", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"username": "abc1",
			"password": "123"
		}`)
		req := _testUtil.ApiRoute("POST", REGISTER_URI, bodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		expected := `{"error":false,"message":"Created user successfully"}`
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 200 Result")

		// execute 2 times
		resp, _ = app.Test(req, -1)
		bodyR, _ = _testUtil.GetBody(resp)
		expected = `{"error":true,"message":"Username is used by another user."}`
		utils.AssertEqual(t, 400, resp.StatusCode, "Return 400 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 400 Result")
	})

	t.Run("[POST] Register - Missing Params", func(t *testing.T) {
		// Missing password
		bodyReader := strings.NewReader(`{
			"username": "abc1",
		}`)
		req := _testUtil.ApiRoute("POST", REGISTER_URI, bodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		expected := `{"error":true,"message":"password: cannot be blank."}`
		utils.AssertEqual(t, 400, resp.StatusCode, "Return 400 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 200 Result")

		// Missing username
		bodyReader = strings.NewReader(`{
			"password": "abc1",
		}`)
		req = _testUtil.ApiRoute("POST", REGISTER_URI, bodyReader, "")
		resp, _ = app.Test(req, -1)
		bodyR, _ = _testUtil.GetBody(resp)
		expected = `{"error":true,"message":"username: cannot be blank."}`
		utils.AssertEqual(t, 400, resp.StatusCode, "Return 400 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 400 Result")
	})
	t.Run("[POST] Login - Successful", func(t *testing.T) {
		// Missing password
		bodyReader := strings.NewReader(`{
			"username": "abc",
			"password": "123"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, bodyReader, "")
		resp, _ := app.Test(req, -1)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")
	})
	t.Run("[POST] Login - Wrong Password", func(t *testing.T) {
		// Missing password
		bodyReader := strings.NewReader(`{
			"username": "abc",
			"password": "1234"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, bodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		expected := `{"error":true,"message":"Invalid credentials"}`
		utils.AssertEqual(t, 401, resp.StatusCode, "Return 401 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 401 Result")
	})
}
