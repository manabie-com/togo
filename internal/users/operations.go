package users

import (
	"log"

	"github.com/kozloz/togo"
)

// UserStore defines the interface needed to store the User resource.
type UserStore interface {
	GetUser(userID int64) (*togo.User, error)
	CreateUser(userID int64) (*togo.User, error)
	UpdateUser(user *togo.User) (*togo.User, error)
}

// Operation is meant to be a reusable class to handle the User resource and logic.
// Attach it to any type of controller(JSON, SOAP, gRPC, etc)
type Operation struct {
	store UserStore
}

// NewOperation is a helper function to create a new Operation
func NewOperation(store UserStore) *Operation {
	return &Operation{
		store: store,
	}
}

// Get gets the User given its ID
func (o *Operation) Get(userID int64) (*togo.User, error) {
	log.Printf("Getting user with ID '%d'.", userID)
	// Get user object
	user, err := o.store.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create creates a user
func (o *Operation) Create(userID int64) (*togo.User, error) {
	log.Printf("Creating user with ID '%d'.", userID)
	// Create user object
	user, err := o.store.CreateUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update updates the user's attributes
func (o *Operation) Update(user *togo.User) (*togo.User, error) {
	log.Printf("Saving user object '%v'.", user)
	// Update user object
	user, err := o.store.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
