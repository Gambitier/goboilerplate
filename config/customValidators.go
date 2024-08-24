package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Add more custom validation fn names here
const (
	portRangeValidation   string = "port_range"
	environmentValidation string = "environment"
)

// Add more custom translations here
var customTranslations = map[string]string{
	portRangeValidation:   "{0} must be in the range of [1024, 65535]",
	environmentValidation: fmt.Sprintf("{0} must be one of [%s, %s, %s]", Development, Staging, Production),
}

// Register custom validation functions here
func registerCustomValidations(validate *validator.Validate) {
	validate.RegisterValidation(portRangeValidation, portRangeValidationFn)
	validate.RegisterValidation(environmentValidation, validateEnvironmentFn)
}

// ========================== Validation logic functions

func portRangeValidationFn(fl validator.FieldLevel) bool {
	port := fl.Field().Uint()
	return port >= 1024 && port <= 65535
}

func validateEnvironmentFn(fl validator.FieldLevel) bool {
	env := fl.Field().Interface().(Environment)
	switch env {
	case Development, Staging, Production:
		return true
	}
	return false
}
