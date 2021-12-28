package api

import (
	"togo/util"

	"github.com/go-playground/validator/v10"
)

var validFullname validator.Func = func(fl validator.FieldLevel) bool {
	if fullname, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedFullname(fullname)
	}
	return false
}
