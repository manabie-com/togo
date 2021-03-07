package src

type (
	IWebServerSetup interface {
		LoadRouterV1() IWebServer
	}

	IWebServer interface {
		Start()
	}
)

type HeaderData struct {
	AccessToken string
}

type IContextService interface {
	LoadContext(data interface{}) error
	CheckPermission(scopes []string) error
	GetTokenData() *TokenData
}

type IErrorFactory interface {
	UnauthorizedError(errorCode string, data error) error
	NotFoundError(errorCode string, data error) error
	InternalServerError(errorCode string, data error) error
	ForbiddenError(errorCode string, data error) error
	BadRequestError(errorCode string, data error) error
}

type TokenData struct {
	UserId      string   `mapstructure:"user_id"`
	Permissions []string `mapstructure:"permissions"`
}

type IJWTService interface {
	CreateToken(data *TokenData) (string, error)
	VerifyToken(token string) (*TokenData, error)
}

const (
	ENTITY_NOT_EXISTS_ERROR = "ENTITY_NOT_EXISTS_ERROR"
	MAPPER_NOT_EXSITS_ERROR = "MAPPER_NOT_EXSITS_ERROR"
	CREATE_TASK_ERROR       = "CREATE_TASK_ERROR"
	FIND_ONE_TASK_ERROR     = "FIND_ONE_TASK_ERROR"
	FIND_TASK_ERROR         = "FIND_TASK_ERROR"
	CREATE_USER_ERROR       = "CREATE_USER_ERROR"
	FIND_USER_ERROR         = "FIND_USER_ERORR"
	FIND_ONE_USER_ERROR     = "FIND_ONE_USER_ERROR"
	USER_IS_NOT_EXISTED     = "USER_IS_NOT_EXISTED"
	CREATE_TOKEN_FAIL       = "CREATE_TOKEN_FAIL"
	TOKEN_INVALID           = "TOKEN_INVALID"
	SERVER_ERROR            = "SERVER_ERROR"
	NO_PERMISSION           = "NO_PERMISSION"
	TOKEN_NOT_PROVIED       = "TOKEN_NOT_PROVIED"
)

const (
	CREATE_TASK = "task.create"
)
