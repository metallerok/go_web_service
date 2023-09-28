package user_services

import (
	"fmt"
	"web_service/src/models"
	"web_service/src/repositories"
)

type UserUpdateDS struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Password string `json:"password"`
}

type IUserUpdaterV2 interface {
	UpdateUserV2(user *models.User, data *UserUpdateDS) (*models.User, error)
}

type UserUpdaterV2 struct {
	UsersRepo repositories.IUsersRepo
}

func (c UserUpdaterV2) UpdateUserV2(user *models.User, data *UserUpdateDS) (*models.User, error) {
	fmt.Println(data)

	return user, nil
}
