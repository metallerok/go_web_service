package webHandlers

import (
	"github.com/gofiber/fiber/v2"
	httpErrors "web_service/src/entrypoints/web/errors"
	services "web_service/src/services/user"
)

func CreateUserAPI(c *fiber.Ctx) error {
	reqUser := &services.UserInputDS{
		Type: "",
		Name: "John",
	}

	err := httpErrors.HTTPValidate(reqUser)

	if err != nil {
		return err
	}

	user := services.CreateUser(reqUser)

	return c.JSON(user)
}
