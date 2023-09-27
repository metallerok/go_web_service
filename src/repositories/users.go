package repositories

import (
	"gorm.io/gorm"
	"web_service/src/models"
)

type IUsersRepo interface {
	Add(user *models.User)
	Get(id int) *models.User
}

type UsersRepo struct {
	DB *gorm.DB
}

func (repo UsersRepo) Add(user *models.User) {
	repo.DB.Create(user)
}

func (repo UsersRepo) Get(id int) *models.User {
	var user models.User
	repo.DB.First(&user, id)

	if user.ID != 0 {
		return &user
	} else {
		return nil
	}
}
