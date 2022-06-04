package response

type ErrorResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type SuccessResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"message,omitempty"`
}
