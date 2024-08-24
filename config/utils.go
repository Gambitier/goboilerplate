package config

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func registerCustomTranslations(validate *validator.Validate, trans ut.Translator) {
	for tag, message := range customTranslations {
		registerCustomTranslation(validate, trans, tag, message)
	}
}

func registerCustomTranslation(validate *validator.Validate, trans ut.Translator, tag, message string) {
	validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
}
