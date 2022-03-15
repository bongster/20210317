.PHONY: all

start:
	@go run src/main.go

go-build:
	@go build -o ./build/app ./src

test:
	@go test -v -cover ./src/...
