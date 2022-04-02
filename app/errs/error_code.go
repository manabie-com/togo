package errs

const (
	ErrCodeInternalAppError = iota + 1
	ErrCodeValidationError
	ErrCodeDbError
	ErrCodeRedisError
	ErrCodeUnauthorized
	ErrCodeTokenIsRequired
	ErrCodeInvalidAuthorizationHeader
	ErrCodeTokenNotFound
	ErrCodeUsernameNotFound
	ErrCodeWrongUserPassword
	ErrCodeCouldNotSaveToken
	ErrCodeReachedLimitDailyTask
)

const (
	ErrInternalAppError           = "internal_app_error"
	ErrValidationError            = "validation_error"
	ErrInternalServiceError       = "internal_service_error"
	ErrTokenIsRequired            = "token_is_required"
	ErrTokenNotFound              = "token_not_found"
	ErrUsernameNotFound           = "username_not_found"
	ErrWrongUserPassword          = "wrong_user_password"
	ErrCouldNotSaveToken          = "could_not_save_token"
	ErrInvalidAuthorizationHeader = "invalid_authorization_header"
	ErrReachedLimitDailyTask      = "reached_limit_daily_task"
)
