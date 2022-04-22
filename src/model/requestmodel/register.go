package requestmodel

// RegisterRequest register request structure
type RegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// LoginRequest login request structure
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
