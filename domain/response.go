package domain

type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Message string `json:"message"`
}
