package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `validate:"user_name" json:"name"`
	Password  string    `json:"-"`
	Type      string    `json:"type"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated"`
	DeletedAt time.Time `json:"-"`
}
