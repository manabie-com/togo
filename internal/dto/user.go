package dto

type LoginRequestDTO struct {
	UserID   string
	Password string
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

type VerifyTokenRequestDTO struct {
	Token string
}

type VerifyTokenResponseDTO struct {
	UserID string
}
