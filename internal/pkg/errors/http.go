package errors

type (
	ValidateError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)
