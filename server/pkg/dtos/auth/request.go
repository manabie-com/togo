package auth

type AuthUserRequest struct {
	Username string `json:"name"`
	Password string `json:"password"`
}
