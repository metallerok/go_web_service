package user_services

import "web_service/src/models"

type UserInputDS struct {
	Name     string `validate:"required,user_name" json:"name"`
	Type     string `validate:"required" json:"type"`
	Password string `validate:"required" json:"password"`
}

func CreateUser(data *UserInputDS) *models.User {
	user := &models.User{
		Id:       1,
		Name:     data.Name,
		Type:     data.Type,
		Password: data.Password,
	}

	return user
}
