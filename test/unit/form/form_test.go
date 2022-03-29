package form_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"togo/models/form"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type BindTestCase struct {
	Input 	map[string]string
	Output 	form.Form
}

var bindTestCases = []BindTestCase{
	{
		Input: map[string]string{"user_id":"", "task_detail":""},
		Output: form.Form{UserID: 0, TaskDetail: ""},
	},
	{
		Input: map[string]string{"user_id":"10", "task_detail":""},
		Output: form.Form{UserID: 10, TaskDetail: ""},
	},
	{
		Input: map[string]string{"user_id":"6", "task_detail":"test binding"},
		Output: form.Form{UserID: 6, TaskDetail: "test binding"},
	},
	{
		Input: map[string]string{"user_id":"", "task_detail":"test again"},
		Output: form.Form{UserID: 0, TaskDetail: "test again"},
	},
}

func TestGinShouldBindForm(t *testing.T) {
	for _, test := range bindTestCases {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, engine := gin.CreateTestContext(w)
		engine.POST("/test-bind", func(c *gin.Context) {
			form := form.Form{}
			c.ShouldBind(&form)
			assert.Equal(t, test.Output, form)
		})

		data := url.Values{}
		data.Add("user_id", test.Input["user_id"])
		data.Add("task_detail", test.Input["task_detail"])
		b := bytes.NewBufferString(data.Encode())

		var err error
		ctx.Request, err = http.NewRequest("POST", "/test-bind", b)
		ctx.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			t.Fatal(err)
		}
		engine.HandleContext(ctx)
	}
}