package validation

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// validate global variable for validator
var validate = validator.New()

// Validate validates data
func Validate(item interface{}) []ValidationError {
	var validationErrors []ValidationErrorsMap

	err := validate.Struct(item)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return validationErrors
		}

		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			tag := strings.ToLower(err.Tag())
			validationError := ValidationErrorsMap[field][tag]
			validationErrors = append(validationErrors, validationError)
		}

	}

	return validationErrors
}
