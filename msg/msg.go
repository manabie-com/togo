package msg

// MsgFlags mapping code and message
var MsgFlags = map[ResponseMessage]string{
	SUCCESS:        "Success",
	ERROR:          "Internal server error",
	INVALID_PARAMS: "Invalid param request",
	UNAUTHORIZED:   "Unauthorized",
	URL_EXPIRED:    "Request has expired",

	ERROR_EXIST:       "ERROR_EXIST",
	ERROR_EXIST_FAIL:  "ERROR_EXIST_FAIL",
	ERROR_NOT_EXIST:   "ERROR_NOT_EXIST",
	ERROR_GET_FAIL:    "Get data failed, please try again",
	ERROR_COUNT_FAIL:  "ERROR_COUNT_FAIL",
	ERROR_ADD_FAIL:    "Add data failed, please try again",
	ERROR_EDIT_FAIL:   "Update data failed, please try again",
	ERROR_DELETE_FAIL: "Delete failed, please try again",
	ERROR_EXPORT_FAIL: "ERROR_EXPORT_FAIL",
	ERROR_IMPORT_FAIL: "ERROR_IMPORT_FAIL",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "ERROR_AUTH_CHECK_TOKEN_FAIL",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "ERROR_AUTH_CHECK_TOKEN_TIMEOUT",
	ERROR_AUTH_TOKEN:               "ERROR_AUTH_TOKEN",
	ERROR_AUTH:                     "ERROR_AUTH",
	ERROR_OKTA_CHECK_TOKEN_TIMEOUT: "Access Token okta has expired or invalid",
	ERROR_AUTH_FAIL:                "The username or password was not correct",
}

//action message

// GetMsg return message with code
func GetMsg(code ResponseMessage) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

func (code ResponseMessage) String() string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func (code ResponseMessage) Error() string {
	return code.String()
}
