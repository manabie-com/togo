package common

type dataBody struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(status int, data interface{}) dataBody {
	return dataBody{Status: status, Data: data, Code: status}
}

func ResponseError(status int, message string) dataBody {
	return dataBody{Status: status, Code: status, Message: message}
}