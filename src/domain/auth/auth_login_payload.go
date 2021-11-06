package auth

// AuthLoginPayload ...
type AuthLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate ...
func (payload AuthLoginPayload) Validate() map[string]string {
	err := make(map[string]string)
	if payload.Username == "" {
		err["message"] = "Invalid Username"
		return err
	}

	if payload.Password == "" {
		err["message"] = "Invalid Password"
		return err
	}

	return nil
}
