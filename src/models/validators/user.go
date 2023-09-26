package validators

import (
	"github.com/go-playground/validator/v10"
)

func ValidateName(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.String() == "" {
		return false
	}

	return true
}
