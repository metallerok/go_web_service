package webapp

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
	"log"
	"time"
	httpErrors "web_service/src/entrypoints/web/errors"
)

func MaskPasswords() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()

		var requestBody map[string]interface{}
		err := json.Unmarshal(body, &requestBody)
		if err != nil {
			return err
		}

		if _, ok := requestBody["password"].(string); ok {
			requestBody["password"] = "********"
		}

		maskedBody, err := json.Marshal(requestBody)
		if err != nil {
			return err
		}

		c.Locals("maskedBody", maskedBody)

		return c.Next()
	}
}

func CustomLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		code := c.Response().StatusCode()

		err := c.Next()

		if err != nil {
			code = c.Response().StatusCode()
		}

		duration := time.Since(start)

		log.Printf("request_id=%s status=%d method=%s path=%s duration=%s body=%s headers=%s\n",
			c.GetRespHeader("X-Request-Id"),
			code,
			c.Method(),
			c.Path(),
			duration.String(),
			c.Locals("maskedBody"),
			c.GetReqHeaders())

		return err
	}
}

func DatabaseMiddleware(dbSession *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("db", dbSession)

		defer func() {
			if r := recover(); r != nil {
				dbSession.Rollback()
			}
		}()

		return c.Next()
	}
}

func MakeApp(dbSession *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: httpErrors.HandleError,
	})

	app.Use(recoverMiddleware.New(recoverMiddleware.Config{
		EnableStackTrace: true,
	}))
	app.Use(DatabaseMiddleware(dbSession))
	app.Use(requestid.New())
	app.Use(MaskPasswords())
	app.Use(CustomLogger())

	SetupRoutes(app)

	return app
}
