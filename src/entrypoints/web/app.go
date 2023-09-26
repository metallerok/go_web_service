package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	httpErrors "web_service/src/entrypoints/web/errors"
	"web_service/src/models/validators"
)

func main() {
	validators.InitValidators()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := fiber.ErrInternalServerError.Message
			internalCode := ""
			var errors_ []string

			var e_ *httpErrors.HTTPError
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
		},
	})

	app.Use(recover.New())
	app.Use(requestid.New())

	requestLogger := logger.New(
		logger.Config{
			Format:     "[${time}]: ${respHeader:X-Request-Id} ${status} - ${method} ${path}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "UTC",
		})
	app.Use(requestLogger)

	SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
