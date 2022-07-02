package response

import (
	"encoding/json"
	"net/http"
	"togo/models"
)

// HandleStatusOK for status ok response
func HandleStatusOK(w http.ResponseWriter, message interface{}, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "Success",
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// HandleStatusCreated for status created response
func HandleStatusCreated(w http.ResponseWriter, message interface{}, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "Success",
		Code:    http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

// HandleStatusBadRequest for status bad request response
func HandleStatusBadRequest(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  "Failed",
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

// HandleStatusInternalServerError for status internal server error response
func HandleStatusInternalServerError(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  "Failed",
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

// HandleStatusNotFound for not found error response
func HandleStatusNotFound(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  "Failed",
		Code:    http.StatusNotFound,
		Message: message,
	})
}
