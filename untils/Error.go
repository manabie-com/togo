package untils

import (
	"TOGO/responses"
	"encoding/json"
	"net/http"
)

func Error(rw http.ResponseWriter, message string, status int) {
	rw.WriteHeader(status)
	response := responses.UserResponse{Status: status, Message: message, Data: map[string]interface{}{"data": nil}}
	json.NewEncoder(rw).Encode(response)
	return
}
