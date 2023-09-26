package httpErrors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"web_service/src/models/validators"
)

type HTTPError struct {
	Code         int      `json:"code"`
	Message      string   `json:"status"`
	InternalCode string   `json:"internal_code"`
	Errors       []string `json:"errors"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewError(code int, internalCode string, errors []string) *HTTPError {
	err := &HTTPError{
		Code:         code,
		Message:      utils.StatusMessage(code),
		InternalCode: internalCode,
		Errors:       errors,
	}
	return err
}

func HTTPValidate(data interface{}) *HTTPError {
	if errs := validators.Validate(data); len(errs) > 0 {

		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, err.Message)
		}

		return NewError(fiber.StatusUnprocessableEntity, "", errMsgs)
	}

	return nil
}
