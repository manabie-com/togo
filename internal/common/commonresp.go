package common

import (
	"encoding/json"
	"net/http"
)

func CommonResponse(resp http.ResponseWriter, code int, obj interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(code)
	json.NewEncoder(resp).Encode(obj)
}
