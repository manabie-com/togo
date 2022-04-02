package auth

import (
	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/server"
)

var ErrEmailExisted = server.NewHTTPValidationError("Email is already existed")
var ErrUsernameExisted = server.NewHTTPValidationError("Username is already existed")

// Register creates new user account
func (a *Auth) Register(data RegisterData) (*model.User, error) {
	// Check if email is already existed
	if rowsAffected := a.db.Take(&model.User{Email: data.Email}).RowsAffected; rowsAffected == 1 {
		return nil, ErrEmailExisted
	}
	// Check if username is already taken yet
	if rowsAffected := a.db.Take(&model.User{Username: data.Username}).RowsAffected; rowsAffected == 1 {
		return nil, ErrUsernameExisted
	}

	u := &model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Username:  data.Username,
		Password:  a.cr.HashPassword(data.Password),
	}

	if err := a.db.Create(u).Error; err != nil {
		return nil, server.NewHTTPInternalError("Error creating user").SetInternal(err)
	}

	return u, nil
}
