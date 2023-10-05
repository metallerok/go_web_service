install:
	go mod download

run_web:
	go run ./src/entrypoints/web

migration:
	atlas migrate diff --env gorm


migrate:
	atlas migrate apply --env gorm


build:
	go build -o ./build/webapp ./src/entrypoints/web/


test:
	go test ./...