package models

// ErrorResponse model for error response
type ErrorResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

// SuccessResponse model for success response
type SuccessResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
