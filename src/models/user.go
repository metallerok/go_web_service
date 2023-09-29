package models

import (
	"time"
)

type UserName string

func (t UserName) Validate() bool {
	if t == "" {
		return false
	}

	return true
}

type UserAge int

func (t UserAge) Validate() bool {
	if t < 0 || t > 200 {
		return false
	}

	return true
}

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      UserName  `validate:"userName" json:"name"`
	Password  string    `json:"-"`
	Type      string    `json:"type"`
	Desc      string    `json:"desc"`
	Age       UserAge   `validate:"userAge" json:"age"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated"`
	DeletedAt time.Time `json:"-"`
}
