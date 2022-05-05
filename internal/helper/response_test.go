package helper

import (
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
)

type HelperTestSuite struct {
	suite.Suite
}

func (suite *HelperTestSuite) TestReturnJSON() {
	var tests = []struct {
		data     map[string]interface{}
		expected string
	}{
		{
			map[string]interface{}{
				"sample_data": "sample",
			},
			`{
				"sample_data": "sample"
			}`,
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		ReturnJSON(w, test.data)
		suite.JSONEq(test.expected, w.Body.String())
	}
}

func TestHelper(t *testing.T) {
	suite.Run(t, new(HelperTestSuite))
}
