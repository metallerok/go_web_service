package user_services

import (
	"github.com/oleiade/reflections"
	"reflect"
	"web_service/src/models"
	"web_service/src/repositories"
)

type UserUpdateDS struct {
	Name     *string `json:"name"`
	Type     *string `json:"type"`
	Password *string `json:"password"`
}

type IUserUpdater interface {
	UpdateUser(user *models.User, data UserUpdateDS) (*models.User, error)
}

type UserUpdater struct {
	UsersRepo repositories.IUsersRepo
}

func (c UserUpdater) UpdateUser(user *models.User, data UserUpdateDS) (*models.User, error) {
	rt := reflect.ValueOf(data)

	if rt.Kind() != reflect.Struct {
		panic("bad UserUpdateDS type")
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Type().Field(i).Name
		value := rt.Field(i)

		if !value.IsNil() {
			err := reflections.SetField(user, field, value.Elem().Interface())

			if err != nil {
				continue
			}
		}

	}

	return user, nil
}
