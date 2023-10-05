package services

import (
	"web_service/src/models"
	"web_service/src/repositories"
)

type UserInputDS struct {
	Name     models.UserName `validate:"required,userName" json:"name"`
	Type     string          `validate:"required" json:"type"`
	Password string          `validate:"required" json:"password"`
	Age      models.UserAge  `validate:"required,userAge" json:"age"`
}

type IUserCreator interface {
	CreateUser(data UserInputDS) models.User
}

type UserCreator struct {
	UsersRepo repositories.IUsersRepo
}

func (c UserCreator) CreateUser(data UserInputDS) models.User {
	user := models.User{
		Name:     data.Name,
		Type:     data.Type,
		Password: data.Password,
		Age:      data.Age,
	}

	c.UsersRepo.Add(&user)

	return user
}
