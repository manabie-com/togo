package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockLoginRequest struct {
	ValidateResponse          error
	GetUserIDResposne         string
	GetPasswordResponse       string
	ShouldBindIsCalled        bool
	ShouldValidateIsCalled    bool
	ShouldGetUserIDIsCalled   bool
	ShouldGetPasswordIsCalled bool
	Testing                   *testing.T
}

func (m MockLoginRequest) Bind(req *http.Request) {
	if !m.ShouldBindIsCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Bind function is called, expected not to be called ")
	}
}
func (m MockLoginRequest) Validate() error {
	if !m.ShouldValidateIsCalled && m.Testing != nil {
		assert.Fail(m.Testing, "Validate function is called, expected not to be called ")
	}
	return m.ValidateResponse
}
func (m MockLoginRequest) GetUserID() string {
	if !m.ShouldGetUserIDIsCalled && m.Testing != nil {
		assert.Fail(m.Testing, "GetUserID function is called, expected not to be called ")
	}
	return m.GetUserIDResposne
}
func (m MockLoginRequest) GetPassword() string {
	if !m.ShouldGetPasswordIsCalled && m.Testing != nil {
		assert.Fail(m.Testing, "GetPassword function is called, expected not to be called ")
	}
	return m.GetPasswordResponse
}
