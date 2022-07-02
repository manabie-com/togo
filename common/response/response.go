package response

import (
	"encoding/json"
	"net/http"
	"togo/models"
)

func HandleStatusOK(w http.ResponseWriter, message interface{}, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "Success",
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func HandleStatusCreated(w http.ResponseWriter, message interface{}, data interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Status:  "Success",
		Code:    http.StatusCreated,
		Message: message,
		Data:    data,
	})
}

func HandleStatusBadRequest(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  "Failed",
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

func HandleStatusInternalServerError(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  "Failed",
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

func HandleStatusNotFound(w http.ResponseWriter, message interface{}) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Status:  "Failed",
		Code:    http.StatusNotFound,
		Message: message,
	})
}
