package models

type User struct {
	Id       int    `json:"id"`
	Name     string `validate:"user_name" json:"name"`
	Password string `json:"-"`
	Type     string `json:"type"`
}
