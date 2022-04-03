package auth

import (
	"net/http"
	"time"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/server"
	"github.com/labstack/echo/v4"
)

var (
	ErrEmailExisted    = server.NewHTTPValidationError("Email is already existed")
	ErrUsernameExisted = server.NewHTTPValidationError("Username is already existed")
	ErrInvalidUsername = server.NewHTTPError(http.StatusBadRequest, "INVALID_USERNAME", "Username is incorrect")
	ErrInvalidEmail    = server.NewHTTPError(http.StatusBadRequest, "INVALID_EMAIL", "Email is incorrect")
	ErrInvalidPassword = server.NewHTTPError(http.StatusBadRequest, "INVALID_PASSWORD", "Password is incorrect")
	ErrGenerateToken   = server.NewHTTPInternalError("Cannot generate token")
)

// Register creates new user account
func (a *Auth) Register(data RegisterData) (*model.User, error) {
	u := &model.User{}
	// Check if email is already existed
	if rowsAffected := a.db.Where(&model.User{Email: data.Email}).Take(u).RowsAffected; rowsAffected == 1 {
		return nil, ErrEmailExisted
	}
	// Check if username is already taken yet
	if rowsAffected := a.db.Where(&model.User{Username: data.Username}).Take(u).RowsAffected; rowsAffected == 1 {
		return nil, ErrUsernameExisted
	}

	u = &model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Username:  data.Username,
		Password:  a.cr.HashPassword(data.Password),
	}

	if err := a.db.Create(u).Error; err != nil {
		return nil, server.NewHTTPInternalError("Cannot creating user").SetInternal(err)
	}

	plan := &model.Plan{}
	if err := a.db.Where(&model.Plan{Name: model.FreemiumPlan}).Take(plan).Error; err != nil {
		return nil, server.NewHTTPError(http.StatusBadRequest, "INVALID_PLAN", "Plan is currently not supported").SetInternal(err)
	}

	// Create default freemium subscription for new user
	sub := &model.Subscription{
		UserID:  u.ID,
		PlanID:  plan.ID,
		StartAt: time.Now(),
		EndAt:   nil,
	}
	if err := a.db.Create(sub).Error; err != nil {
		return nil, server.NewHTTPInternalError("Cannot create subscription").SetInternal(err)
	}

	return u, nil
}

// LoginUsername logins user using his/her username
func (a *Auth) LoginUsername(data LoginUsernameData) (*model.AuthToken, error) {
	user := &model.User{}

	// Check if username is in system
	if err := a.db.Where(&model.User{Username: data.Username}).Take(user).Error; err != nil {
		return nil, ErrInvalidUsername
	}

	// Check if password is matched
	if !a.cr.CompareHashAndPassword(user.Password, data.Password) {
		return nil, ErrInvalidPassword
	}

	// Create new JWT for successfully logging in user
	claims := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}

	tokenStr, duration, err := a.jwt.GenerateToken(claims)
	if err != nil {
		return nil, ErrGenerateToken.SetInternal(err)
	}

	token := &model.AuthToken{
		AccessToken: tokenStr,
		ExpiresIn:   duration,
	}

	return token, nil
}

// LoginEmail logins user using his/her email
func (a *Auth) LoginEmail(data LoginEmailData) (*model.AuthToken, error) {
	user := &model.User{}

	// Check if username is in system
	if err := a.db.Where(&model.User{Email: data.Email}).Take(user).Error; err != nil {
		return nil, ErrInvalidEmail
	}

	// Check if password is matched
	if !a.cr.CompareHashAndPassword(user.Password, data.Password) {
		return nil, ErrInvalidPassword
	}

	// Create new JWT for successfully logging in user
	claims := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}

	tokenStr, duration, err := a.jwt.GenerateToken(claims)
	if err != nil {
		return nil, ErrGenerateToken.SetInternal(err)
	}

	token := &model.AuthToken{
		AccessToken: tokenStr,
		ExpiresIn:   duration,
	}

	return token, nil
}

// User returns user data stored in jwt token
func (s *Auth) User(c echo.Context) *model.AuthUser {
	id, _ := c.Get("id").(float64)
	user, _ := c.Get("username").(string)
	email, _ := c.Get("email").(string)

	return &model.AuthUser{
		ID:       int(id),
		Username: user,
		Email:    email,
	}
}
