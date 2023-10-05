package webapp

import (
	"github.com/gofiber/fiber/v2"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
	httpErrors "web_service/src/entrypoints/web/errors"
)

func MakeApp(dbSession *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: httpErrors.HandleError,
	})

	app.Use(recoverMiddleware.New(recoverMiddleware.Config{
		EnableStackTrace: true,
	}))
	app.Use(DatabaseMiddleware(dbSession))
	app.Use(requestid.New())
	app.Use(MaskSecrets())
	app.Use(CustomLogger())

	SetupRoutes(app)

	return app
}
