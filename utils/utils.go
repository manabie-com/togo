package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type response struct {
	Status  string
	Message string
	Data    interface{}
}

func SuccessRespond(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	res := response{
		Status:  "Success",
		Message: message,
		Data:    data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}
func FailureRespond(w http.ResponseWriter, statusCode int, message string) {
	res := response{
		Status:  "Failure",
		Message: message,
		Data:    nil,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func Str2Uint32(str string) (uint32, error) {
	u64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u64), nil
}
