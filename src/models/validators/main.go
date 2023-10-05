package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
	"reflect"
	"strings"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
	Message     string
}

var Validator_ = validator.New()

func InitValidators() {
	Validator_.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := Validator_.RegisterValidation("userName", ValidateUserName); err != nil {
		log.Fatal(err)
	}

	if err := Validator_.RegisterValidation("userAge", ValidateUserAge); err != nil {
		log.Fatal(err)
	}
}

func Validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := Validator_.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Message = msgForTag(err.Tag())

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func msgForTag(tag string) string {
	switch tag {
	case "email":
		return "invalid email"
	case "userAge":
		return "age must be > 0 and < 200"
	case "userName":
		return "wrong name"
	}
	return tag
}
