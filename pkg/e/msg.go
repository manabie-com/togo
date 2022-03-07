package e

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := Msg[code]
	if ok {
		return msg
	}
	return Msg[ERROR]
}

var Msg = map[int]string{
	SUCCESS: "success",
	ERR:     "error",

	SUCCESS_LOGIN:      "success_login",
	ERROR:              "Internal Server Error",
	ERROR_UNAUTHORIZES: "Unauthorized",
	INVALID_PARAMS:     "Request parameter error",

	ERROR_AUTH_CHECK_TOKEN_FAIL:            "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:         "Token已超时",
	ERROR_AUTH:                             "Invalid Username Or Password",
	ERROR_IN_EXCEED_LIMIT_TASK_ADD_PER_DAY: "You had exceeded the limit task added per day. Please try again later.",
}
