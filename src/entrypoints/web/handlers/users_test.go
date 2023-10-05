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
)

const userUrl = "/user"

func TestCreateUserAPI(t *testing.T) {
	dbSession := db.MakeDB()
	app := webapp.MakeApp(dbSession)

	reqBody_ := map[string]interface{}{
		"name":     "Test name",
		"password": "1234",
		"type":     "admin",
		"age":      18,
	}
	reqBody, _ := json.Marshal(reqBody_)

	req := httptest.NewRequest("POST", userUrl, bytes.NewReader(reqBody))

	resp, _ := app.Test(req)

	if resp.StatusCode != fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Errorf(
			"CreateUserAPI error for '%s' url with status code: %d and body: %s",
			userUrl, resp.StatusCode, body,
		)
	}
}
