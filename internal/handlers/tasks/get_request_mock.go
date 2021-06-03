package tasks

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockGetRequest struct {
	BindResponse          error
	ValidateResponse      error
	GetCreateDateResponse string

	ShouldBindCalled          bool
	ShouldValidateCalled      bool
	ShouldGetCreateDateCalled bool
	Testing                   *testing.T
}

func (m MockGetRequest) Bind(*http.Request) error {
	if !m.ShouldBindCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Bind function is called, expected not to be called")
	}
	return m.BindResponse
}

func (m MockGetRequest) Validate() error {
	if !m.ShouldValidateCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Validate function is called, expected not to be called")
	}
	return m.ValidateResponse
}

func (m MockGetRequest) GetCreateDate() string {
	if !m.ShouldGetCreateDateCalled && m.Testing != nil {
		assert.Fail(m.Testing, "GetCreateDate function is called, expected not to be called")
	}
	return m.GetCreateDateResponse
}
