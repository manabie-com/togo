package repository

import (
	"togo/models"
)

// Define `User` Repository Interface with the following
// Methods which will be utilized by the `UserService`
type UserRepository interface {
	// Add new users to the database
	Register(user *models.User) (*models.User, error)

	// Check if user exists in the database
	GetUserByEmail(email string, user *models.User) (models.User, error)

	// Get user by JWT token, to identify if session exists
	GetUserByToken(token string) (models.User, error)

	// Update token to maintain session
	Login(user *models.User) error
}
