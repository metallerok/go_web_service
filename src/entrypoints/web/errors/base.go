package httpErrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/utils"
	"web_service/src/models/validators"
)

func HandleError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := fiber.ErrInternalServerError.Message
	internalCode := ""
	var errors_ map[string]string

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

	log.Error(err)

	return ctx.Status(code).JSON(fiber.Map{
		"status":        message,
		"code":          code,
		"internal_code": internalCode,
		"errors":        errors_,
	})
}

func HandleMarshalingError(err error) error {
	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		errMsg := make(map[string]string)
		errMsg[typeErr.Field] = fmt.Sprintf(
			"Wrong field type: %s. Must be %s",
			typeErr.Value,
			typeErr.Type.Kind())

		return NewError(
			fiber.StatusUnprocessableEntity, "", errMsg)
	}

	return err
}

type HTTPError struct {
	Code         int               `json:"code"`
	Message      string            `json:"status"`
	InternalCode string            `json:"internal_code"`
	Errors       map[string]string `json:"errors"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewError(code int, internalCode string, errors map[string]string) *HTTPError {
	err := &HTTPError{
		Code:         code,
		Message:      utils.StatusMessage(code),
		InternalCode: internalCode,
		Errors:       errors,
	}
	return err
}

func HTTPValidate(data interface{}) error {
	errs := validators.Validate(data)
	if errs == nil || len(errs) == 0 {
		return nil
	}

	errMsgs := make(map[string]string)

	for _, err := range errs {
		errMsgs[err.FailedField] = err.Message
	}

	return NewError(fiber.StatusUnprocessableEntity, "", errMsgs)
}
