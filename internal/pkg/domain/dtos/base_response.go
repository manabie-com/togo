package dtos

// BaseResponse struct
type BaseResponse struct {
	Status int            `json:"status"`
	Data   interface{}    `json:"data"`
	Error  *ErrorResponse `json:"error"`
}

// ErrorResponse struct
type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
