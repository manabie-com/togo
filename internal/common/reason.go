package common

import (
	"encoding/json"
	"net/http"
)

type ReasonCode string

const (
	ReasonUnauthorized             ReasonCode = "-1"
	ReasonCreateTokenError         ReasonCode = "-2"
	ReasonUserIDPasswordEmptyError ReasonCode = "-3"
	ReasonInvalidToken             ReasonCode = "-4"
	ReasonInternalError            ReasonCode = "-5"
	ReasonInvalidArgument          ReasonCode = "-6"
	ReasonNotFound                 ReasonCode = "-7"
	ReasonDateInvalidFormat        ReasonCode = "-8"
	ReasonExceededLimit            ReasonCode = "-9"
)

var reasonCodeValues = map[string]string{
	"-1": "incorrect user_id/pwd",
	"-2": "create token error",
	"-3": "user_id/pwd must not empty",
	"-4": "invalid token",
	"-5": "server internal error",
	"-6": "invalid argument",
	"-7": "data not found",
	"-8": "invalid format yyyy-mm-dd",
	"-9": "task exceeded limit per day",
}

func (rc ReasonCode) Code() string {
	return string(rc)
}

func (rc ReasonCode) Message() string {
	if value, ok := reasonCodeValues[rc.Code()]; ok {
		return value
	}
	return "general error"
}

func parseError(err error) ReasonCode {
	return ReasonCode(err.Error())
}

func ResponseError(err error, resp http.ResponseWriter) {
	switch parseError(err).Code() {
	case ReasonUnauthorized.Code():
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": parseError(err).Message(),
		})
	case ReasonCreateTokenError.Code(),
		ReasonInternalError.Code():
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": parseError(err).Message(),
		})
	case ReasonUserIDPasswordEmptyError.Code(),
		ReasonInvalidArgument.Code(),
		ReasonDateInvalidFormat.Code():
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": parseError(err).Message(),
		})
	case ReasonInvalidToken.Code():
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": parseError(err).Message(),
		})
	case ReasonNotFound.Code():
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": parseError(err).Message(),
		})
	case ReasonExceededLimit.Code():
		resp.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": parseError(err).Message(),
		})
	}
}
