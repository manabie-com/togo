package user

import (
	"togo/src/domain/user"
	"togo/src/schema"
)

type UserController struct {
	workflow user.IUserWorkflow
}

func (this *UserController) Login(data *schema.LoginRequest) (*schema.LoginResponse, error) {
	return this.workflow.Login(data)
}

func (this *UserController) Register(data *schema.RegisterRequest) (*schema.RegisterResponse, error) {
	return this.workflow.Register(data)
}

func (this *UserController) CreateTaskByOwner(data *schema.CreateTaskByOwnerRequest) (*schema.CreateTaskByOwnerResponse, error) {
	return nil, nil
}

func (this *UserController) DeleteTaskByOwner(data *schema.DeleteTaskByOwnerRequest) (*schema.DeleteTaskByOwnerResponse, error) {
	return nil, nil
}

func NewUserController() IUserController {
	return &UserController{
		user.NewUserWorkflow(),
	}
}
