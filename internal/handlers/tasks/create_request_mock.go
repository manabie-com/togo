package tasks

import (
	"net/http"
	"testing"

	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockCreateRequest struct {
	BindResponse     error
	ValidateResponse error
	ToModelResponse  *models.Task

	ShouldBindCalled     bool
	ShouldValidateCalled bool
	ShouldToModelCalled  bool
	Testing              *testing.T
}

func (m MockCreateRequest) Bind(*http.Request) error {
	if !m.ShouldBindCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Bind function is called, expected not to be called")
	}
	return m.BindResponse
}

func (m MockCreateRequest) Validate() error {
	if !m.ShouldValidateCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Validate function is called, expected not to be called")
	}
	return m.ValidateResponse
}

func (m MockCreateRequest) ToModel(string) *models.Task {
	if !m.ShouldToModelCalled && m.Testing != nil {
		assert.Fail(m.Testing, "ToModel function is called, expected not to be called")
	}
	return m.ToModelResponse
}
