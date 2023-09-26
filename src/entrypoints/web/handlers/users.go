package webHandlers

import (
	"github.com/gofiber/fiber/v2"
	httpErrors "web_service/src/entrypoints/web/errors"
	services "web_service/src/services/user"
)

func CreateUserAPI(c *fiber.Ctx) error {
	userInput := new(services.UserInputDS)

	if err := c.BodyParser(userInput); err != nil {
		return err
	}

	err := httpErrors.HTTPValidate(userInput)

	if err != nil {
		return err
	}

	user := services.CreateUser(userInput)

	return c.JSON(user)
}
