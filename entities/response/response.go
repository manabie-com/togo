package response

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) Error(err error) Response {
	r.Status = false
	r.Message = err.Error()
	r.Data = nil

	return *r
}

func (r *Response) Success(data interface{}) Response {
	r.Status = true
	r.Message = "Successful operation"
	r.Data = data

	return *r
}
