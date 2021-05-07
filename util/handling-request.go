package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Abort(c *gin.Context, httpStatusCode int, errorCode, detail interface{}) {
	resp := gin.H{
		"status_code": httpStatusCode,
		"error_code":  errorCode,
		"detail":      detail,
	}

	c.AbortWithStatusJSON(httpStatusCode, resp)
}

func AbortUnauthorized(c *gin.Context, errorCode, detail interface{}) {
	Abort(c, http.StatusUnauthorized, errorCode, detail)
}

func AbortEntityNotFound(c *gin.Context, errorCode, detail interface{}) {
	Abort(c, http.StatusServiceUnavailable, errorCode, detail)
}

func AbortUnexpected(c *gin.Context, errorCode, detail interface{}) {
	Abort(c, http.StatusInternalServerError, errorCode, detail)
}

func AbortJSONBadRequest(c *gin.Context) {
	Abort(c, http.StatusBadRequest, ERR_CODE_JSON_UNMARSHAL, "Can't parse data")
}

func AbortAlreadyExists(c *gin.Context, errorCode, detail interface{}) {
	Abort(c, http.StatusConflict, errorCode, detail)
}

const (
	ERR_CODE_JSON_UNMARSHAL = "JSON_UNMARSHAL"
	ERR_CODE_DB_ISSUE       = "ISSUE"

	ERR_CODE_DECRYPT_JWT_TOKEN                = "CAN'T_DECRYPT_TOKEN"
	ERR_CODE_NOT_EVEN_A_TOKEN                 = "NOT_EVEN_A_TOKEN"
	ERR_CODE_PERMISSION_DENY                  = "PERMISSION_DENY"
	ERR_CODE_TOKEN_EXPIRED                    = "TOKEN_EXPIRED"
	ERR_CODE_COULD_NOT_HANDLE_THIS_TOKEN      = "COULD_NOT_HANDLE_THIS_TOKEN"
	ERR_CODE_INVALID_TOKEN                    = "INVALID_TOKEN"
	ERR_CODE_AUTH_WRONG_USER_NAME_OR_PASSWORD = "AUTH_WRONG_USER_NAME_OR_PASSWORD"

	ERR_CODE_USER_EXISTED          = "USER_EXISTED"
	ERR_CODE_USER_REACH_LIMIT_TASK = "USER_REACH_LIMIT_TASK"
)
