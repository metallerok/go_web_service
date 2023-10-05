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
