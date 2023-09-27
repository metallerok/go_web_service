package repositories

import (
	"gorm.io/gorm"
	"web_service/src/models"
)

type IUsersRepo interface {
	Add(user *models.User)
}

type UsersRepo struct {
	DB *gorm.DB
}

func (repo UsersRepo) Add(user *models.User) {
	repo.DB.Create(user)
}
