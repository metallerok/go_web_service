package main

import (
	"log"
	"web_service/src/db"
	webapp "web_service/src/entrypoints/web/app"
	"web_service/src/models/validators"
)

func main() {
	validators.InitValidators()

	dbSession := db.MakeDB()
	app := webapp.MakeApp(dbSession)

	log.Fatal(app.Listen(":8000"))
}
