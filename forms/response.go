package forms

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
