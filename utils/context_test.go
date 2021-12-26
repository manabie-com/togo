package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExtractFromConText(t *testing.T) {
	id := "someID"
	ctx := AddToContext(context.Background(), id)

	result, ok := ExtractFromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, id, result)
}
