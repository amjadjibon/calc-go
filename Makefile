.PHONY: all build

build:
	@go build -o bin/calc ./cmd/calc/main.go

run: build
	@./bin/calc

lint:
	@golangci-lint run --fix --timeout 5m
test:
	@go test -v ./...

cover:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

release: build
	@goreleaser release --rm-dist

clean:
	@rm -rf dist
	@rm -f coverage.out coverage.html
