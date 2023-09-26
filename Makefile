install:
	go mod download

run_web:
	go run src/entrypoints/web/app.go