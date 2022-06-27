package constants

const (
	FileEnvironment = ".env"
	DsnEnvironment  = "DSN_POSTGRES"
)

// Response
const (
	ResponseStatus = "status"

	// Error
	ResponseError       = "error"
	ResponseErrorDetail = "error_detail"

	// Ok
	ResponseMessage = "message"
	ResponseData    = "data"
)

// Key Cookie
const (
	CookieTokenKey string = "token"
)
