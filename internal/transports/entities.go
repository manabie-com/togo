package transports

type LoginRequest struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
