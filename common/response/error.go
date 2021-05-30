package response

// ErrorResponse follows JSON error response format
type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}
