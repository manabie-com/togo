package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

var (
	translator   ut.Translator
	translations = map[string]string{
		"required_with": "{0} is required with {1}",
	}
	validations = map[string]func(fl validator.FieldLevel) bool{
		"date": DateValidation,
	}
)

// NewValidator /**
func NewValidator() (*validator.Validate, error) {

	v := validator.New()
	// name func registration
	registerTagNameFunc(v)
	// translations registration
	if err := registerTranslations(v); err != nil {
		return nil, err
	}
	// validations registration
	if err := registerValidations(v); err != nil {
		return nil, err
	}

	return v, nil
}

// Translate /**
func Translate(err error) map[string][]string {
	result := make(map[string][]string)

	errs := err.(validator.ValidationErrors)

	for _, err := range errs {
		result[err.Field()] = append(result[err.Field()], err.Translate(translator))
	}

	return result
}

// registerTranslations /**
func registerTranslations(v *validator.Validate) error {

	// default local is english
	enLocal := en.New()
	uni := ut.New(enLocal)
	translator, _ = uni.GetTranslator("en")
	// register translations
	if err := enTranslations.RegisterDefaultTranslations(v, translator); err != nil {
		return err
	}

	// register custom translations
	for tag, message := range translations {
		registerFn := func(ut ut.Translator) error {
			return ut.Add(tag, message, false)
		}
		translateFn := func(ut ut.Translator, fe validator.FieldError) string {
			param := fe.Param()
			tag := fe.Tag()
			t, err := ut.T(tag, fe.Field(), param)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		}
		if err := v.RegisterTranslation(tag, translator, registerFn, translateFn); err != nil {
			return err
		}
	}

	return nil
}

// registerTagNameFunc /**
func registerTagNameFunc(v *validator.Validate) {
	// register tag name func
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// registerValidations /**
func registerValidations(v *validator.Validate) error {
	// register custom validations
	for tag, customFunc := range validations {
		if err := v.RegisterValidation(tag, customFunc); err != nil {
			return err
		}
	}

	return nil
}
