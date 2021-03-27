package servehttp

type AppHandler struct {
	Route       string
	Handler     IAPIHandler
	Method      string
	// More options for register handler add here. eg: timeout, middleware...
}
