package validate_test

import (
	"testing"
	"togo/globals/validator"
	"togo/models/form"

	"github.com/stretchr/testify/assert"
)

type ValidateTestCase struct {
	Input form.Form
	Output []validator.ErrorJSON
}

var validateTestCases = []ValidateTestCase{
	{
		Input: form.Form{UserID: 0, TaskDetail: ""},
		Output: []validator.ErrorJSON{
			{Namespace: "Form.UserID", ActualTag: "required"},
			{Namespace: "Form.TaskDetail", ActualTag: "required"},
		},
	},
	{
		Input: form.Form{UserID: 1, TaskDetail: ""},
		Output: []validator.ErrorJSON{
			{Namespace: "Form.TaskDetail", ActualTag: "required"},
		},
	},
	{
		Input: form.Form{UserID: 0, TaskDetail: "test some cases"},
		Output: []validator.ErrorJSON{
			{Namespace: "Form.UserID", ActualTag: "required"},
		},
	},
	{
		Input: form.Form{UserID: 10, TaskDetail: "this case should pass"},
		Output: nil,
	},
}

func TestFormShouldBeValidated(t *testing.T) {
	for _,test := range validateTestCases {
		res := validator.Validate(test.Input)
		assert.Equal(t, test.Output, res)
	}
}