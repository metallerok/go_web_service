package webapp

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"time"
	httpErrors "web_service/src/entrypoints/web/errors"
)

func MaskSecrets() fiber.Handler {
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
			code = fiber.StatusInternalServerError

			var e_ *httpErrors.HTTPError
			if errors.As(err, &e_) {
				code = e_.Code
			}

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
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
