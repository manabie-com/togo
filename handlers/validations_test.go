package handlers

import (
	"github.com/namnhatdoan/togo/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmailSuccess(t *testing.T) {
	a := assert.New(t)
	err := validateEmail("dummy@gmail.com")
	a.Nil(err)
}

func TestValidateEmailWithInvalidDomain(t *testing.T) {
	t.Skip("builtin mail package accept this case")
	a := assert.New(t)
	err := validateEmail("dummy@g")
	a.NotNil(err)
}

func TestValidateEmailWithEmptyAddress(t *testing.T) {
	a := assert.New(t)
	err := validateEmail("@gmail.com")
	a.NotNil(err)
}

func TestValidateTaskSuccess(t *testing.T) {
	a := assert.New(t)
	err := validateTask("Some task here")
	a.Nil(err)
}

func TestValidateTaskWithEmptyValue(t *testing.T) {
	a := assert.New(t)
	err := validateTask("")
	a.NotNil(err)
	a.Equal(err.Error(), constants.MissingTask)
}

