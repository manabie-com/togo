package model_test

import (
	"fmt"
	"github.com/manabie-com/togo/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

// hash password
func TestHashPassword(t *testing.T) {
	fmt.Println("pwdHash: ", model.HashPassword("123456"))
	assert.True(t, false)
}
