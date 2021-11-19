package services

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// ConvertJSONValueToVariable func
func ConvertJSONValueToVariable(FieldName string, JSONObject map[string]interface{}) (string, interface{}) {
	valPostFormObject := JSONObject[FieldName]
	valPostForm := fmt.Sprintf("%v", valPostFormObject)
	// @TODO - if valPostFormObject == nil => no pass fiedl , if valPostFormObject <> nil => pass fiedl
	if valPostFormObject == nil {
		valPostFormObjectLower := JSONObject[strings.ToLower(FieldName)]
		valPostFormLower := fmt.Sprintf("%v", valPostFormObjectLower)
		return valPostFormLower, valPostFormObjectLower
	}
	// has pass field
	return valPostForm, valPostFormObject
}

func GenerateUserID() string {
	return uuid.New().String()
}
