package utils

const (
	SuccessRequestMessage = "success"
	UnauthorizedRequestMessage = "Unauthorized"
)

type Meta struct {
	Code int
	Message string
}

func BuildSuccessResponseRequest(meta *Meta, data interface{}) map[string]interface{}{
	resp := make(map[string]interface{})
	resp["meta"] = &meta
	resp["data"] = data
	return resp
}

func BuildErrorResponseRequest(meta *Meta) map[string]interface{}{
	resp := make(map[string]interface{})
	resp["meta"] = &meta
	return resp
}