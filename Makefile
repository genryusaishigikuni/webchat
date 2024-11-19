build:
	@go build -o bin/webchat chat_service/main.go

test:
	@go test -v ./...

run: build
	@./bin/webchat

migration:
	@migrate create -ext sql -dir chat_service/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run chat_service/migrate/main.go up

migrate-down:
	@go run chat_service/migrate/main.go down


