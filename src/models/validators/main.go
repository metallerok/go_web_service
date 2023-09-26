package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
	Message     string
}

var Validator_ *validator.Validate = validator.New()

func InitValidators() {
	err := Validator_.RegisterValidation("user_name", ValidateName)

	if err != nil {
		log.Fatal(err)
	}
}

func Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := Validator_.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Message = fmt.Sprintf(
				"%s validation error with value '%s' by reason %s",
				elem.FailedField,
				elem.Value,
				elem.Tag,
			)

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
