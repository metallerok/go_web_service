package validators

import (
	"github.com/go-playground/validator/v10"
	"web_service/src/models"
)

func ValidateUserName(fl validator.FieldLevel) bool {
	f := fl.Field().Interface().(models.UserName)
	return f.Validate()
}

func ValidateUserAge(fl validator.FieldLevel) bool {
	f := fl.Field().Interface().(models.UserAge)
	return f.Validate()
}
