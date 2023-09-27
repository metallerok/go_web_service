package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

		code := c.Response().StatusCode()

		err := c.Next()

		if err != nil {
			code = fiber.StatusInternalServerError
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

func DatabaseMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		dsn := "host=localhost user=datagrip password=datagrip dbname=test_gorm port=5432 sslmode=disable TimeZone=UTC"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Panicf(err.Error())
		}
		c.Locals("db", db)

		return c.Next()
	}
}

func main() {
	validators.InitValidators()

	app := fiber.New(fiber.Config{
		ErrorHandler: httpErrors.HandleError,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(DatabaseMiddleware())
	app.Use(requestid.New())
	app.Use(MaskPasswords())
	app.Use(CustomLogger())

	SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
