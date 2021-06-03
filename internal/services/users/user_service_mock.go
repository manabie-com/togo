package users

import (
	"context"
	"testing"

	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	ValidateResponse         error
	GetUserByIDResponseUser  *models.User
	GetUserByIDResponseError error

	ShouldValidateIsCalled  bool
	ShouldGetUserByIDCalled bool
	Testing                 *testing.T
}

func (m MockUserService) Validate(context.Context, string, string) error {
	if !m.ShouldValidateIsCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Validate function is called, expected not to be called")
	}
	return m.ValidateResponse
}

func (m MockUserService) GetUserByID(context.Context, string) (*models.User, error) {
	if !m.ShouldGetUserByIDCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Validate function is called, expected not to be called")
	}
	return m.GetUserByIDResponseUser, m.GetUserByIDResponseError
}
