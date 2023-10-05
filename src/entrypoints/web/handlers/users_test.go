package webHandlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http/httptest"
	"testing"
	"web_service/src/db"
	webapp "web_service/src/entrypoints/web/app"
	"web_service/src/repositories"
)

const userUrl = "/user"

func TestCreateUserAPI(t *testing.T) {
	DB := db.MakeDB()
	app := webapp.MakeApp(DB)

	reqBody_ := map[string]interface{}{
		"name":     "Test name",
		"password": "1234",
		"type":     "admin",
		"age":      18,
	}
	reqBody, _ := json.Marshal(reqBody_)

	req := httptest.NewRequest("POST", userUrl, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf(
			"CreateUserAPI error for '%s' with status code: %d and resp.body: %s",
			userUrl, resp.StatusCode, body,
		)
	}

	var respBody map[string]interface{}
	if err = json.Unmarshal(body, &respBody); err != nil {
		t.Error(err)
	}

	createUserId := int(respBody["id"].(float64))

	if createUserId == 0 {
		t.Errorf("CreateUserAPI error: user was not saved")
	}
	if respBody["name"] != reqBody_["name"] {
		t.Errorf("CreateUserAPI error: wrong user.name = %s. Must be %s", respBody["name"], reqBody_["name"])
	}
	if respBody["type"] != reqBody_["type"] {
		t.Errorf("CreateUserAPI error: wrong user.type = %s. Must be %s", respBody["type"], reqBody_["type"])
	}

	respBodyAge := int(respBody["age"].(float64))
	if respBodyAge != reqBody_["age"] {
		t.Errorf("CreateUserAPI error: wrong user.age = %d. Must be %d", respBody["age"], reqBody_["age"])
	}

	usersRepo := repositories.UsersRepo{DB: DB}
	dbUser, err := usersRepo.Get(createUserId)

	if err != nil {
		t.Error(err)
	}

	if dbUser.ID != uint(createUserId) {
		t.Errorf("CreateUserAPI error: fetched user_id != respBody['id']")
	}
}

func TestCreateUserAPIButWrongSchemaType(t *testing.T) {
	DB := db.MakeDB()
	app := webapp.MakeApp(DB)

	reqBody_ := map[string]interface{}{
		"name":     "Test name",
		"password": "1234",
		"type":     "admin",
		"age":      18,
	}
	reqBody, _ := json.Marshal(reqBody_)

	req := httptest.NewRequest("POST", userUrl, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/text")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusUnprocessableEntity {
		t.Errorf("CreateUserAPI test error: wrong schema type must raise 422")
	}
}

func TestCreateUserAPIButWrongSchemaFieldType(t *testing.T) {
	DB := db.MakeDB()
	app := webapp.MakeApp(DB)

	reqBody_ := map[string]interface{}{
		"name":     "Test name",
		"password": 1234,
		"type":     "admin",
		"age":      18,
	}
	reqBody, _ := json.Marshal(reqBody_)

	req := httptest.NewRequest("POST", userUrl, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusUnprocessableEntity {
		t.Errorf("CreateUserAPI test error: wrong schema must raise 422")
	}

	var respBody map[string]interface{}
	if err = json.Unmarshal(body, &respBody); err != nil {
		t.Error(err)
	}

	respErrors := respBody["errors"].(map[string]interface{})

	if _, ok := respErrors["password"]; !ok {
		t.Errorf("CreateUserAPI test error: wrong password type must be in error messages")
	}
}

func TestCreateUserAPIButNotValidSchema(t *testing.T) {
	DB := db.MakeDB()
	app := webapp.MakeApp(DB)

	reqBody_ := map[string]interface{}{
		"name":     "",
		"password": "1234",
		"type":     "admin",
		"age":      -1,
	}
	reqBody, _ := json.Marshal(reqBody_)

	req := httptest.NewRequest("POST", userUrl, bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != fiber.StatusUnprocessableEntity {
		t.Errorf("CreateUserAPI test error: validation must raise 422")
	}

	var respBody map[string]interface{}
	if err = json.Unmarshal(body, &respBody); err != nil {
		t.Error(err)
	}

	respErrors := respBody["errors"].(map[string]interface{})

	if _, ok := respErrors["name"]; !ok {
		t.Errorf("CreateUserAPI test error: name validation error must be in error messages")
	}
	if _, ok := respErrors["age"]; !ok {
		t.Errorf("CreateUserAPI test error: age validation error must be in error messages")
	}
}
