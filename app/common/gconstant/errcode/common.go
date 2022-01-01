package errcode

const (
	ServerErrorCommon             = "SERVER_ERR_COMMON"
	UserErrCommon                 = "USER_ERR_COMMON"
	TokenIsRequired               = "TOKEN_IS_REQUIRED"
	TokenIsInvalid                = "TOKEN_IS_INVALID"
	IPIsInvalid                   = "IP_IS_INVALID"
	AgentIsInvalid                = "AGENT_IS_INVALID"
	OTPNotValid                   = "OTP_NOT_VALID"
	OTPSendFail                   = "OTP_SEND_FAIL"
	StatusNotSupported            = "STATUS_NOT_SUPPORTED"
	TokenIsExpired                = "TOKEN_IS_EXPIRED"
	APIKeyOrAPISecretKeyIsInvalid = "API_KEY_OR_API_SECRET_KEY_IS_INVALID"
	TimeIsInvalid                 = "TIME_IS_INVALID"
)
