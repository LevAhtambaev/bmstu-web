dbUp:
	docker compose up -d

build:
	go build -o bin/main cmd/WAD-2022/main.go

migrate:
	go run cmd/migrate/main.go

run:
	go run cmd/WAD-2022/main.go