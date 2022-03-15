package functional_test

import (
	"encoding/json"
	"strings"
	"testing"
	_testUtil "togo-service/tests/utils"

	"github.com/gofiber/fiber/v2/utils"
)

const SETTING_URI = "/api/v1/admin/user-setting"

func Test_Setting(t *testing.T) {
	t.Run("[POST] Admin can do setting", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"username": "admin",
			"password": "secret"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, bodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var adminRes _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &adminRes)

		bodyReader = strings.NewReader(`{
			"user_id": 2,
			"quota_per_day": 3
		}`)
		req = _testUtil.ApiRoute("POST", SETTING_URI, bodyReader, adminRes.Token)
		resp, err := app.Test(req, -1)
		if err != nil {
			println(err.Error())
		}
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")
	})

	t.Run("[POST] User cannot do setting", func(t *testing.T) {
		bodyReader := strings.NewReader(`{
			"username": "user",
			"password": "secret"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, bodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var logres _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &logres)

		bodyReader = strings.NewReader(`{
			"user_id": 2,
			"quota_per_day": 3
		}`)
		req = _testUtil.ApiRoute("POST", SETTING_URI, bodyReader, logres.Token)
		resp, err := app.Test(req, -1)
		if err != nil {
			println(err.Error())
		}
		utils.AssertEqual(t, 403, resp.StatusCode, "Return 403 Code")
	})
}
