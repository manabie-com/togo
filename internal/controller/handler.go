package controller

type Handler struct {
	UserController
	TaskController
	AuthController
}

var HandlerInstance *Handler

func NewHandler(userController UserController, taskController TaskController, authController AuthController) *Handler {
	return &Handler{
		UserController: userController,
		TaskController: taskController,
		AuthController: authController,
	}
}
