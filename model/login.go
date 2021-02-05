package model

type (
	LoginInput struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Data  string `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}
)
