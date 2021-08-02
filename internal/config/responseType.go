package config

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response with JSON format
// swagger:model Response
type responser struct {
	// Status
	Success bool `json:"success"`
	// Message
	Message string `json:"message,omitempty"`
	// Payload
	Data interface{} `json:"data,omitempty"`
}

// Response with JSON format
// swagger:response jsonResponse
type jsonResponse struct {
	Body responser `json:"ErrorLists"`
}

// ResponseWithJSON formats returned data
func ResponseWithJSON(w http.ResponseWriter, success bool, msg string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	response := responser{success, msg, payload}

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(&response)
}
func ResponseWithErrorJSON(w http.ResponseWriter, success bool, msg string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	response := responser{success, msg, payload}

	encoder := json.NewEncoder(w)
	_ = encoder.Encode(&response)
}

/**
@desc response HTML content
@param w ResponseWriter, response string
@return avoid
*/
func ResponseWithHtml(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}

// ResponseWithError returns an error
func ResponseWithError(w http.ResponseWriter, msg string, err error) {
	if err != nil {
		log.Printf("The detailed error: %v \n", err)
	}
	ResponseWithErrorJSON(w, false, msg, nil)
}

// ResponseWithError returns an error
func ResponseWithDetailedError(w http.ResponseWriter, msg string, err error, details interface{}) {
	if err != nil {
		log.Printf("The detailed error: %v \n", err)
	}
	ResponseWithErrorJSON(w, false, msg, details)
}

// ResponseWithSuccess returns an error
func ResponseWithSuccess(w http.ResponseWriter, msg string, payload interface{}) {
	ResponseWithJSON(w, true, msg, payload)
}

func ResponseSucessAndLog(w http.ResponseWriter, msg string, payload interface{}, request string, handler string) {
	// logger.Entry().WithFields(logger.Fields{
	// 	"handler": handler,
	// 	"request": request,
	// }).Info(msg)
	ResponseWithJSON(w, true, msg, payload)
}
