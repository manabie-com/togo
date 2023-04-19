// Package authenticate authenticates a user
package authenticate

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// Set of error variables for the feature.
var (
	ErrNotFound              = errors.New("user not found")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	GetUserByEmail(ctx context.Context, email mail.Address) (User, error)
}

// Feature manages the api for create todo.
type Feature struct {
	storer Storer
}

// NewFeature constructs a feature for create todo api access.
func NewFeature(storer Storer) *Feature {
	return &Feature{
		storer: storer,
	}
}

// Create inserts a new user into the database.
func (f *Feature) Authenticate(ctx context.Context, email mail.Address, password string) (User, error) {
	usr, err := f.storer.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	if err := bcrypt.CompareHashAndPassword(usr.PasswordHash, []byte(password)); err != nil {
		return User{}, fmt.Errorf("comparehashandpassword: %w", ErrAuthenticationFailure)
	}

	return usr, nil
}
