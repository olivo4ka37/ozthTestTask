.PHONY:
.DEFAULT_GOAL := run
build:
		go mod download
		go build -o  ./.bin/app cmd/main.go

run:	build compose migrate
		./.bin/app

migrate:
		migrate -path ./migrations -database postgres://postgres:qwerty@localhost:5030/techdb?sslmode=disable up

dropTables:
		migrate -path ./migrations -database postgres://postgres:qwerty@localhost:5030/techdb?sslmode=disable down

compose:
		docker-compose up -d

