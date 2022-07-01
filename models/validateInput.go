package models

import (
	"strings"
)

// Check user input is valid or not
func CheckUserInput(user NewUser) bool {
	password := strings.TrimSpace(user.Password)
	username := strings.TrimSpace(user.Username)
	if password == "" || username == "" {
		return false
	}
	return true
}

// Check task input value is valid or not
func CheckTaskInput(task NewTask) bool {
	var Content string
	if task.Content != "" {
		Content = strings.TrimSpace(task.Content)
	}
	if Content == "" {
		return false
	}
	return true
}
