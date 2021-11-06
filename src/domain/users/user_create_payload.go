package users

// UserCreatePayload ...
type UserCreatePayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate ...
func (payload UserCreatePayload) Validate() map[string]string {
	err := make(map[string]string)

	if payload.Username == "" {
		err["message"] = "Invalid Username"
		return err
	}

	if len(payload.Username) < 6 || len(payload.Username) > 50 {
		err["message"] = "Username must greater 6 and less than 50 "
		return err
	}

	if payload.Password == "" {
		err["message"] = "Invalid Password"
		return err
	}

	if len(payload.Password) < 6 || len(payload.Password) > 50 {
		err["message"] = "Password must greater 6 and less than 50 "
		return err
	}

	return nil
}
