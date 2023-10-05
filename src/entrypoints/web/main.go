package main

import (
	"log"
	"web_service/src/db"
	webapp "web_service/src/entrypoints/web/app"
)

func main() {
	dbSession := db.MakeDB()
	app := webapp.MakeApp(dbSession)

	log.Fatal(app.Listen(":8000"))
}
