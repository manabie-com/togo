package common_test

import (
	"errors"
	"github.com/japananh/togo/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUID_String(t *testing.T) {
	var tcs = []struct {
		arg      common.UID
		expected string
	}{
		{common.NewUID(1, 1, 1), "e532qos8jjM2"},
		{common.NewUID(2, 2, 1), "gGzTCUsW2BnY"},
	}

	for _, tc := range tcs {
		assert.EqualValues(t, tc.arg.String(), tc.expected)
	}
}

func TestUID_GetLocalID(t *testing.T) {
	var tcs = []struct {
		arg      common.UID
		expected uint32
	}{
		{common.NewUID(1, 1, 1), 1},
		{common.NewUID(2, 2, 1), 2},
	}

	for _, tc := range tcs {
		assert.EqualValues(t, tc.arg.GetLocalID(), tc.expected)
	}
}

func TestUID_GetObjectType(t *testing.T) {
	var tcs = []struct {
		arg      common.UID
		expected int
	}{
		{common.NewUID(1, 1, 1), 1},
		{common.NewUID(2, 2, 1), 2},
	}

	for _, tc := range tcs {
		assert.EqualValues(t, tc.arg.GetObjectType(), tc.expected)
	}
}

func TestUID_GetShardID(t *testing.T) {
	var tcs = []struct {
		arg      common.UID
		expected uint32
	}{
		{common.NewUID(1, 1, 1), 1},
		{common.NewUID(2, 2, 3), 3},
	}

	for _, tc := range tcs {
		assert.EqualValues(t, tc.arg.GetShardID(), tc.expected)
	}
}

func TestUID_NewUID(t *testing.T) {
	var tcs = []struct {
		localID    uint32
		objectType int
		shardID    uint32
		expected   string
	}{
		{1, 1, 1, "e532qos8jjM2"},
		{2, 1, 1, "gGzTBURqhajG"},
	}

	for _, tc := range tcs {
		uid := common.NewUID(tc.localID, tc.objectType, tc.shardID)
		assert.EqualValues(t, tc.expected, uid.String())
	}
}

func TestUID_FromBase58(t *testing.T) {
	var tcs = []struct {
		arg      string
		expected common.UID
		err      error
	}{
		{"e532qos8jjM2", common.NewUID(1, 1, 1), nil},
		{"e532qos8jj", common.UID{}, errors.New("wrong uid")},
	}

	for _, tc := range tcs {
		uid, err := common.FromBase58(tc.arg)
		assert.ObjectsAreEqualValues(tc.expected, uid)
		if tc.err != nil {
			assert.NotNil(t, err)
		}
	}
}
