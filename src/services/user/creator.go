package user_services

import "web_service/src/models"

type UserInputDS struct {
	Name string `validate:"required,user_name"`
	Type string `validate:"required"`
}

func CreateUser(data *UserInputDS) *models.User {
	user := &models.User{
		Id:   1,
		Name: data.Name,
		Type: data.Type,
	}

	return user
}
