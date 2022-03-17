.PHONY: all

start:
	@go run src/main.go

dev:
	./bin/air

go-build:
	@go build -o ./build/app ./src

test:
	@go test -v -cover ./src/...
