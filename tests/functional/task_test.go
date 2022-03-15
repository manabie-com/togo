package functional_test

import (
	"encoding/json"
	"strings"
	"testing"
	_testUtil "togo-service/tests/utils"

	"github.com/gofiber/fiber/v2/utils"
)

const CREATE_TASK_URI = "/api/v1/task"
const GET_TASK_URI = "/api/v1/tasks"
const UPDATE_TASK_URI = "/api/v1/task/1"

func Test_Task(t *testing.T) {
	t.Run("[POST] User can create task", func(t *testing.T) {
		// guest cannot
		req := _testUtil.ApiRoute("POST", CREATE_TASK_URI, nil, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		expected := `{"error":true,"msg":"Missing or malformed JWT"}`

		utils.AssertEqual(t, 400, resp.StatusCode, "Return 400 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 400 Result")

		// user login
		loginBodyReader := strings.NewReader(`{
			"username": "user",
			"password": "secret"
		}`)
		req = _testUtil.ApiRoute("POST", LOGIN_URI, loginBodyReader, "")
		resp, _ = app.Test(req, -1)
		bodyR, _ = _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var logres _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &logres)

		// user create task w/o params
		newreq := _testUtil.ApiRoute("POST", CREATE_TASK_URI, nil, logres.Token)
		resp, _ = app.Test(newreq, -1)
		bodyR, _ = _testUtil.GetBody(resp)

		expected = `{"error":true,"message":"description: cannot be blank; name: cannot be blank."}`
		utils.AssertEqual(t, 400, resp.StatusCode, "Return 400 Code")
		utils.AssertEqual(t, expected, string(bodyR), "Return 200 Result")

		// user create task successfully
		taskBodyReader := strings.NewReader(`{
			"name": "here is taskname",
			"description": "some desc"
		}`)

		newreq = _testUtil.ApiRoute("POST", CREATE_TASK_URI, taskBodyReader, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 201, resp.StatusCode, "Return 201 Code")
	})

	t.Run("[POST] User cannot create task over quota", func(t *testing.T) {
		resetDB(db)
		// user login
		loginBodyReader := strings.NewReader(`{
			"username": "user",
			"password": "secret"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, loginBodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var logres _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &logres)

		taskBodyReader := strings.NewReader(`{
			"name": "here is taskname1",
			"description": "some desc1"
		}`)

		// created 1st task
		newreq := _testUtil.ApiRoute("POST", CREATE_TASK_URI, taskBodyReader, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 201, resp.StatusCode, "Return 201 Code")

		taskBodyReader = strings.NewReader(`{
			"name": "here is taskname2",
			"description": "some desc2"
		}`)

		// created 2nd task
		newreq = _testUtil.ApiRoute("POST", CREATE_TASK_URI, taskBodyReader, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 201, resp.StatusCode, "Return 201 Code")

		// failed to create 3rd task
		taskBodyReader = strings.NewReader(`{
			"name": "here is taskname3",
			"description": "some desc3"
		}`)
		newreq = _testUtil.ApiRoute("POST", CREATE_TASK_URI, taskBodyReader, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 422, resp.StatusCode, "Return 422 Code")
	})
	t.Run("[GET] Get all user task", func(t *testing.T) {
		// user login
		loginBodyReader := strings.NewReader(`{
			"username": "user",
			"password": "secret"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, loginBodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var logres _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &logres)

		newreq := _testUtil.ApiRoute("GET", GET_TASK_URI, nil, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")
	})

	t.Run("User can update - delete own task", func(t *testing.T) {
		resetDB(db)
		// user1 login
		loginBodyReader := strings.NewReader(`{
			"username": "user",
			"password": "secret"
		}`)
		req := _testUtil.ApiRoute("POST", LOGIN_URI, loginBodyReader, "")
		resp, _ := app.Test(req, -1)
		bodyR, _ := _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var logres _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &logres)

		// user2 login
		loginBodyReader = strings.NewReader(`{
			"username": "user2",
			"password": "secret"
		}`)
		req = _testUtil.ApiRoute("POST", LOGIN_URI, loginBodyReader, "")
		resp, _ = app.Test(req, -1)
		bodyR, _ = _testUtil.GetBody(resp)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		var logres2 _testUtil.LoginRes
		json.Unmarshal([]byte(bodyR), &logres2)

		taskBodyReader := strings.NewReader(`{
			"name": "here is taskname1",
			"description": "some desc1"
		}`)

		// created 1st task
		newreq := _testUtil.ApiRoute("POST", CREATE_TASK_URI, taskBodyReader, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 201, resp.StatusCode, "Return 201 Code")

		taskBodyReader = strings.NewReader(`{
			"name": "here is taskname new",
			"description": "some desc new"
		}`)
		// user 1 can update task
		newreq = _testUtil.ApiRoute("PUT", UPDATE_TASK_URI, taskBodyReader, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")

		taskBodyReader = strings.NewReader(`{
			"name": "here is taskname new",
			"description": "some desc new"
		}`)
		// user 2 cannot update user1's task
		newreq = _testUtil.ApiRoute("PUT", UPDATE_TASK_URI, taskBodyReader, logres2.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 422, resp.StatusCode, "Return 422 Code")

		// user 2 cannot delete user1's task
		newreq = _testUtil.ApiRoute("DELETE", UPDATE_TASK_URI, nil, logres2.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 422, resp.StatusCode, "Return 422 Code")

		// user 1 can delete own task
		newreq = _testUtil.ApiRoute("DELETE", UPDATE_TASK_URI, nil, logres.Token)
		resp, _ = app.Test(newreq, -1)
		utils.AssertEqual(t, 200, resp.StatusCode, "Return 200 Code")
	})
}
