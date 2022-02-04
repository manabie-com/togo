package users

import "github.com/kozloz/togo"

type UserStore interface {
	Get(userID int64) (*togo.User, error)
}
type Operation struct {
	store UserStore
}

// Create the task for the user
func (o *Operation) Get(userID int64) (*togo.User, error) {
	// Get user object

	user, err := o.store.Get(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
