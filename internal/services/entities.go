package services

type ApiDataResp struct {
	Data interface{} `json:"data"`
}

func newDataResp(data interface{}) *ApiDataResp {
	return &ApiDataResp{Data: data}
}

type ApiErrResp struct {
	Error interface{} `json:"error"`
}

func newErrResp(err interface{}) *ApiErrResp {
	return &ApiErrResp{Error: err}
}