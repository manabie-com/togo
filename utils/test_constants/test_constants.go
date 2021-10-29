package test_constants

const (
	UserName               = "firstUser"
	Password               = "example"
	CreatedDate            = "2021-10-24"
	TestJWTKey             = "wqGyEBBfPK9w3Lxw"
	LoginUrl               = "/login?user_id=" + UserName + "&password=" + Password
	GetTasksUrl            = "/tasks?created_date=" + CreatedDate
	PostTasksUrl           = "/tasks"
	HeaderAuthorizationKey = "Authorization"
	InvalidToken           = "INVALID_TOKEN"
	TaskContent            = "test content"
	CreateTaskReqBody      = `{"content":"` + TaskContent + `"}`
	DefaultMaxToDo         = 5
	WrongUserName          = "wrongUserName"
	WrongPassword          = "wrongPassword"
)
