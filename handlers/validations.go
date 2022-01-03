package handlers

import (
	"errors"
	"github.com/namnhatdoan/togo/constants"
	"net/mail"
)

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func validateTask(task string) error {
	if task == "" {
		return errors.New(constants.MissingTask)
	}
	return nil
}