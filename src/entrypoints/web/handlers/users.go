package webHandlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	httpErrors "web_service/src/entrypoints/web/errors"
	"web_service/src/repositories"
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

	db, ok := c.Locals("db").(*gorm.DB)

	if !ok {
		return fmt.Errorf("failed to get database connection from context")
	}

	var usersRepo repositories.IUsersRepo = &repositories.UsersRepo{
		DB: db,
	}

	var userCreator services.IUserCreator = &services.UserCreator{
		UsersRepo: usersRepo,
	}

	user := userCreator.CreateUser(userInput)

	db.Commit()

	return c.JSON(user)
}

func GetUserAPI(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")

	if err != nil {
		return httpErrors.NewError(fiber.StatusUnprocessableEntity, "", make([]string, 0))
	}

	db, ok := c.Locals("db").(*gorm.DB)

	if !ok {
		return fmt.Errorf("failed to get database connection from context")
	}

	var usersRepo repositories.IUsersRepo = &repositories.UsersRepo{
		DB: db,
	}

	user := usersRepo.Get(id)

	if user == nil {
		return httpErrors.NewError(fiber.StatusNotFound, "", make([]string, 0))
	}

	return c.JSON(user)
}

func UpdateUserAPI(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return httpErrors.NewError(fiber.StatusUnprocessableEntity, "", make([]string, 0))
	}

	reqBody := new(services.UserUpdateDS)

	if err := c.BodyParser(reqBody); err != nil {
		return err
	}

	err = httpErrors.HTTPValidate(reqBody)

	if err != nil {
		return err
	}

	db, ok := c.Locals("db").(*gorm.DB)

	if !ok {
		return fmt.Errorf("failed to get database connection from context")
	}

	var usersRepo repositories.IUsersRepo = &repositories.UsersRepo{
		DB: db,
	}

	user := usersRepo.Get(id)

	if user == nil {
		return httpErrors.NewError(fiber.StatusNotFound, "", make([]string, 0))
	}

	var userUpdater services.IUserUpdater = &services.UserUpdater{
		UsersRepo: usersRepo,
	}

	user, err = userUpdater.UpdateUser(user, reqBody)

	if err != nil {
		return httpErrors.NewError(fiber.StatusBadRequest, "", make([]string, 0))
	}

	db.Save(user)
	db.Commit()

	return c.JSON(user)
}
