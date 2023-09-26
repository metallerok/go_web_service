install:
	go mod download

run_web:
	go run src/entrypoints/web/app.go

migration:
	atlas migrate diff --env gorm


migrate:
	atlas schema apply --env gorm