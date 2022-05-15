package common_test

import (
	"github.com/japananh/togo/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUID_GenUID(t *testing.T) {
	var tcs = []struct {
		id       int
		dbType   int
		expected string
	}{
		{1, 1, "e532qos8jjM2"},
		{2, 1, "gGzTBURqhajG"},
	}

	for _, tc := range tcs {
		sqlModel := common.SQLModel{Id: tc.id}
		sqlModel.GenUID(tc.dbType)
		assert.EqualValues(t, tc.expected, sqlModel.FakeId.String())

	}
}
