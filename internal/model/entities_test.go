package model_test

import (
	"fmt"
	"testing"

	"github.com/manabie-com/togo/internal/model"
	"github.com/stretchr/testify/assert"
)

// hash password
func TestHashPassword(t *testing.T) {
	fmt.Println("pwdHash: ", model.HashPassword("example"))
	assert.True(t, false)
}
