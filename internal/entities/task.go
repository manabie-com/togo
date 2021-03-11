package entities

import (
	"regexp"
	"strings"
)

// Task reflect Task entity in general which is decoupled from
// database entity
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
}

// TaskValidationError represents all validation errors of task
type TaskValidationError string

func (e TaskValidationError) Error() string {
	return string(e)
}

var (
	// ErrTaskInvalidCreatedDate ...
	ErrTaskInvalidCreatedDate = TaskValidationError("Invalid created date")
	// ErrTaskInvalidContent ...
	ErrTaskInvalidContent = TaskValidationError("Content of task must not be empty")
)

// ValidateCreatedDate ensure that createdDate must be defined in format dd-mm-yyyy
func (t *Task) ValidateCreatedDate() bool {
	re := regexp.MustCompile("^((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$")
	return re.MatchString(t.CreatedDate)
}

// NormalizeContent ...
func (t *Task) NormalizeContent() error {
	t.Content = strings.Trim(t.Content, " ")
	ok := t.ValidateContent()
	if !ok {
		return ErrTaskInvalidContent
	}
	return nil
}

// ValidateContent ...
func (t *Task) ValidateContent() bool {
	if len(t.Content) == 0 {
		return false
	}
	return true
}
