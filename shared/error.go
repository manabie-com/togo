package shared

import (
	"errors"
	"fmt"
	apperror "github.com/manabie-com/togo/shared/app_error"
	"strings"
)

var ErrNotFound = apperror.NewCustomError(errors.New("record not found"), "record not found", "ErrNotFound")

func ErrInternal(err error) *apperror.AppError {
	return apperror.NewErrorResponse(err, err.Error(), "ErrInternal")
}

func ErrInvalidRequest(err error) *apperror.AppError {
	return apperror.NewErrorResponse(err, err.Error(), "ErrInvalidRequest")
}

func ErrCannotGetEntity(entity string, err error) *apperror.AppError {
	return apperror.NewCustomError(
		err,
		fmt.Sprintf("Cannot Get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *apperror.AppError {
	return apperror.NewCustomError(
		err,
		fmt.Sprintf("Cannot Delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *apperror.AppError {
	return apperror.NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *apperror.AppError {
	return apperror.NewCustomError(
		err,
		fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}
