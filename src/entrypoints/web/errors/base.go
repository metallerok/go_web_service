package httpErrors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"web_service/src/models/validators"
)

func HandleError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := fiber.ErrInternalServerError.Message
	internalCode := ""
	var errors_ []string

	var e_ *HTTPError
	if errors.As(err, &e_) {
		internalCode = e_.InternalCode
		code = e_.Code
		message = e_.Message
		errors_ = e_.Errors
	}

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = utils.StatusMessage(e.Code)
	}

	return ctx.Status(code).JSON(fiber.Map{
		"status":        message,
		"code":          code,
		"internal_code": internalCode,
		"errors":        errors_,
	})
}

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
