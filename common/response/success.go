package response

// SuccessResponse follows JSON success response format
type SuccessResponse struct {
	Status  int         `json:"status" form:"status"`
	Data    interface{} `json:"data" form:"data"`
	Message string      `json:"message" form:"message"`
}
