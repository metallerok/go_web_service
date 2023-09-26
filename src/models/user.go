package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `validate:"user_name" json:"name"`
	Password  string    `json:"-"`
	Type      string    `json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
