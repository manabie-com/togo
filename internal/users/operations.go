package users

import (
	"log"

	"github.com/kozloz/togo"
)

type UserStore interface {
	GetUser(userID int64) (*togo.User, error)
	CreateUser() (*togo.User, error)
	UpdateUser(user *togo.User) (*togo.User, error)
}
type Operation struct {
	store UserStore
}

func NewOperation(store UserStore) *Operation {
	return &Operation{
		store: store,
	}
}

// Get gets the User given its ID
func (o *Operation) Get(userID int64) (*togo.User, error) {
	// Get user object

	user, err := o.store.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Create creates a user
func (o *Operation) Create() (*togo.User, error) {
	// Create user object
	user, err := o.store.CreateUser()
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
