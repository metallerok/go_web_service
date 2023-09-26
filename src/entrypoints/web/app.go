package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"time"
	httpErrors "web_service/src/entrypoints/web/errors"
	"web_service/src/models/validators"
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

		err := c.Next()

		duration := time.Since(start)

		log.Printf("status=%d method=%s path=%s duration=%s body=%s headers=%s\n",
			c.Response().StatusCode(),
			c.Method(),
			c.Path(),
			duration.String(),
			c.Locals("maskedBody"),
			c.GetReqHeaders())

		return err
	}
}

func main() {
	validators.InitValidators()

	app := fiber.New(fiber.Config{
		ErrorHandler: httpErrors.HandleError,
	})

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(MaskPasswords())
	app.Use(CustomLogger())

	SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
