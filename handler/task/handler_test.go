package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewHandler(t *testing.T) {
	mockRecordService := new(mockRecordService)
	handler := NewHandler(mockRecordService)
	assert.IsType(t, &Handler{}, handler)
}