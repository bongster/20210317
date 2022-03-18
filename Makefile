.PHONY: all

check-air:
ifeq (,$(wildcard ./bin/air))
	curl -sSfL "https://raw.githubusercontent.com/cosmtrek/air/master/install.sh" | sh -s
endif

start:
	@go run src/main.go

dev: check-air
	./bin/air

go-build:
	@go build -o ./build/app ./src

test:
	@go test -v -cover ./src/...
