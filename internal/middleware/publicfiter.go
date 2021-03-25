package middleware

import "net/http"

type PublicFilter struct {
	filter PublicHandler
}

type PublicHandler struct {
	handler HandlerFunc
}

// DoFilter do filter then handle the api controller
func (pf *PublicHandler) DoFilter(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	pf.handler(resp, req)
}

func NewPublicHandler(wrapper HandlerFunc) *PublicHandler {
	return &PublicHandler{handler: wrapper}
}
