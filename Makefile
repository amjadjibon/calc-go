.PHONY: all build

build:
	@go build -o bin/calc ./cmd/calc/main.go

run: build
	@./bin/calc

lint:
	@golangci-lint run --fix --timeout 5m
test:
	@go test -v ./...
