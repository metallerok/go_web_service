package user_services

import (
	"web_service/src/models"
	"web_service/src/repositories"
)

type UserInputDS struct {
	Name     string `validate:"required,user_name" json:"name"`
	Type     string `validate:"required" json:"type"`
	Password string `validate:"required" json:"password"`
}

type IUserCreator interface {
	CreateUser(data *UserInputDS) *models.User
}

type UserCreator struct {
	UsersRepo repositories.IUsersRepo
}

func (c UserCreator) CreateUser(data *UserInputDS) *models.User {
	user := &models.User{
		Name:     data.Name,
		Type:     data.Type,
		Password: data.Password,
	}

	c.UsersRepo.Add(user)

	return user
}
