package domain

import (
	"context"
	"errors"
)

var (
	// ErrUnauthorized unauthorized request error
	ErrUnauthorized = errors.New("UNAUTHORIZED")
	// ErrLoginFailed failed to login error
	ErrLoginFailed = errors.New("LOGIN_FAILED")
	// ErrCredentialInvalid invalid login request error
	ErrCredentialInvalid = errors.New("CREDENTIAL_INVALID")
)

// LoginCredential parameters
type LoginCredential struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

// LoginResult result of successfully login
type LoginResult struct {
	Profile *User  `json:"profile"`
	Token   string `json:"token"`
}

// VerifyTokenResult Verify and decode token result
type VerifyTokenResult struct {
	Authenticated bool  `json:"authenticated"`
	Payload       *User `json:"payload"`
}

// AuthService auth service interface
type AuthService interface {
	Login(ctx context.Context, credential *LoginCredential) (*LoginResult, error)
	VerifyToken(ctx context.Context, token string) (*VerifyTokenResult, error)
}
