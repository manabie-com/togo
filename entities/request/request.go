package request

type UserRequest struct {
	Username string ` json:"username" validate:"required,gte=0,lte=256" `
	Password string ` json:"password" validate:"required"`
}
