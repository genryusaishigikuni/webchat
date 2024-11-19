build:
	@go build -o bin/webchat chat_service/main.go

test:
	@go test -v ./...

run: build
	@./bin/webchat