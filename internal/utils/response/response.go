package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status  string
	Message string
	Code    int
	Data    interface{}
	Date    time.Time
}

func Render(w http.ResponseWriter, data interface{}, code int, message string, status string) {
	json.NewEncoder(w).Encode(Response{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
		Date:    time.Now(),
	})
}
