package repositories

import (
	"gorm.io/gorm"
	"web_service/src/models"
)

type IUsersRepo interface {
	Add(user *models.User)
	Get(id int) (*models.User, error)
}

type UsersRepo struct {
	DB *gorm.DB
}

func (repo UsersRepo) Add(user *models.User) {
	repo.DB.Create(user)
}

func (repo UsersRepo) Get(id int) (*models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
