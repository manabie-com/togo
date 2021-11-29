package msg

type ResponseMessage int

const (
	SUCCESS                        ResponseMessage = 200
	ERROR                          ResponseMessage = 500
	INVALID_PARAMS                 ResponseMessage = 400
	UNAUTHORIZED                   ResponseMessage = 401
	URL_EXPIRED                    ResponseMessage = 410
	ERROR_OKTA_CHECK_TOKEN_TIMEOUT ResponseMessage = 415
	ERROR_EXIST                    ResponseMessage = 1001
	ERROR_EXIST_FAIL               ResponseMessage = 1002
	ERROR_NOT_EXIST                ResponseMessage = 1003
	ERROR_GET_FAIL                 ResponseMessage = 1004
	ERROR_COUNT_FAIL               ResponseMessage = 1005
	ERROR_ADD_FAIL                 ResponseMessage = 1006
	ERROR_EDIT_FAIL                ResponseMessage = 1007
	ERROR_DELETE_FAIL              ResponseMessage = 1008
	ERROR_EXPORT_FAIL              ResponseMessage = 1009
	ERROR_IMPORT_FAIL              ResponseMessage = 1010

	ERROR_AUTH_FAIL                ResponseMessage = 2001
	ERROR_AUTH_CHECK_TOKEN_FAIL    ResponseMessage = 2002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT ResponseMessage = 2003
	ERROR_AUTH_TOKEN               ResponseMessage = 2004
	ERROR_AUTH                     ResponseMessage = 2005
)
