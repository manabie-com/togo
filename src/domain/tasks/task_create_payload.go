package tasks

// TaskCreatePayload ...
type TaskCreatePayload struct {
	Content string `json:"content"`
}

// Validate ...
func (payload TaskCreatePayload) Validate() map[string]string {
	err := make(map[string]string)

	if payload.Content == "" {
		err["message"] = "Invalid Content"
		return err
	}

	return nil
}
