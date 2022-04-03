package model

// AuthToken contains authorized token details
type AuthToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID       int
	Username string
	Email    string
}
