package user_services

import (
	"github.com/go-errors/errors"
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
	"web_service/src/models"
	"web_service/src/repositories"
)

//type UserUpdateDS struct {
//	Name     string `validate:"required,user_name" json:"name"`
//	Type     string `validate:"required" json:"type"`
//	Password string `validate:"required" json:"password"`
//}

type IUserUpdater interface {
	UpdateUser(user *models.User, data *map[string]interface{}) (*models.User, error)
}

type UserUpdater struct {
	UsersRepo repositories.IUsersRepo
}

func buildFieldsByTagMap(key string, s interface{}) map[string]string {
	var fieldsByTag = make(map[string]string)

	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		panic("bad type")
	}

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get(key), ",")[0] // use split to ignore tag "options"
		if v == "" || v == "-" {
			continue
		}
		fieldsByTag[v] = f.Name
	}

	return fieldsByTag
}

func (c UserUpdater) UpdateUser(user *models.User, data *map[string]interface{}) (*models.User, error) {
	userFields := buildFieldsByTagMap("json", *user)
	for key, value := range *data {
		if key == "password" {
			err := reflections.SetField(user, "Password", value)

			if err != nil {
				return nil, errors.Wrap(err, 0)
			}
		}

		if fieldName, ok := userFields[key]; ok {
			err := reflections.SetField(user, fieldName, value)

			if err != nil {
				return nil, errors.Wrap(err, 0)
			}
		}
	}

	return user, nil
}
