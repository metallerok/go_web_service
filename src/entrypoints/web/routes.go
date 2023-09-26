package main

import (
	"github.com/gofiber/fiber/v2"
	webHandlers "web_service/src/entrypoints/web/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", webHandlers.CreateUserAPI)
}
