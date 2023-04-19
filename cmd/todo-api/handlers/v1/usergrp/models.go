package usergrp

// =============================================================================

// AppAuthenticatingUser contains information needed to authenticate a user.
type AppAuthenticatingUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
