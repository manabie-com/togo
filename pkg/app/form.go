package app

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate
var uni *ut.UniversalTranslator

// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func init() {
	// NOTE: ommitting allot of error checking for brevity
	en := en.New()
	uni = ut.New(en, en)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func BindAndValid(c *gin.Context, form interface{}) error {
	err := c.ShouldBindBodyWith(form, binding.JSON)
	if err != nil {
		return err
	}
	er := validate.Struct(form)
	if er != nil {
		return er
	}
	return nil
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) string {
	switch reflect.TypeOf(err).Kind().String() {
	case reflect.Ptr.String():
		return err.(error).Error()
	case reflect.Slice.String():
		trans, _ := uni.GetTranslator("en")
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			return e.Translate(trans)
		}
	}
	return ""
}
