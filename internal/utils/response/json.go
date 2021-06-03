package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func JSON(resp http.ResponseWriter, data interface{}) error {
	resp.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(resp).Encode(Response{
		Data: data,
	}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Error(resp http.ResponseWriter, err error) error {
	if decodeErr := json.NewEncoder(resp).Encode(Response{
		Error: err.Error(),
	}); decodeErr != nil {
		log.Println(decodeErr)
	}

	return err
}

func InternalError(resp http.ResponseWriter, err error) error {
	resp.WriteHeader(http.StatusInternalServerError)
	return Error(resp, err)
}
func NotFoundPath(resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusNotFound)
}
