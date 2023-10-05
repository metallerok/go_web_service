package services

import (
	"testing"
	"web_service/src/db"
	"web_service/src/repositories"
)

func TestCreateUser(t *testing.T) {
	data := UserInputDS{
		Name:     "Test name",
		Type:     "Admin",
		Password: "1234",
		Age:      18,
	}

	dbSession := db.MakeDB()

	usersRepo := repositories.UsersRepo{
		DB: dbSession,
	}

	userCreator := UserCreator{
		UsersRepo: usersRepo,
	}

	user, err := userCreator.CreateUser(data)

	if err != nil {
		t.Error(err)
	}

	dbSession.Commit()

	if user.Name != data.Name {
		t.Errorf("UserCreator error: user.Name = %s. Must be %s", user.Name, data.Name)
	}
	if user.Type != data.Type {
		t.Errorf("UserCreator error: user.Type = %s. Must be %s", user.Type, data.Type)
	}
	if user.Password != data.Password {
		t.Errorf("UserCreator error: user.Password = %s. Must be %s", user.Password, data.Password)
	}
	if user.Age != data.Age {
		t.Errorf("UserCreator error: user.Age = %d. Must be %d", user.Age, data.Age)
	}
	if user.CreatedAt.IsZero() {
		t.Errorf("UserCreator error: user.CreatedAt must not be Zero time")
	}
	if user.UpdatedAt.IsZero() {
		t.Errorf("UserCreator error: user.UpdatedAt must not be Zero time")
	}

	if time, _ := user.DeletedAt.Value(); time != nil {
		t.Errorf("UserCreator error: user.DeletedAt must be Zero time")
	}
}
