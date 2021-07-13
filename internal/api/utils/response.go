package utils

import (
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/api/dictionary"
	"github.com/manabie-com/togo/internal/pkg/logger"
	"net/http"
)

type CommonResponse struct {
	Data interface{} `json:"data"`
}

type ErrorCommonResponse struct {
	ErrorStr interface{} `json:"error"`
}

// WriteJSON create response with data or error if exist
func WriteJSON(ctx context.Context, rsp http.ResponseWriter, statusCode int, data interface{}, err error) {
	rsp.Header().Set("Content-Type", "application/json")
	rsp.WriteHeader(statusCode)

	if err != nil {
		if jsonErr := json.NewEncoder(rsp).Encode(ErrorCommonResponse{ErrorStr: err.Error()}); jsonErr != nil {
			logger.MBErrorf(ctx, "%s: %s", dictionary.FailedToEncodeJSON, jsonErr)
		}
		return
	}

	if jsonErr := json.NewEncoder(rsp).Encode(CommonResponse{Data: data}); jsonErr != nil {
		logger.MBErrorf(ctx, "%s: %s", dictionary.FailedToEncodeJSON, jsonErr)
	}
}
