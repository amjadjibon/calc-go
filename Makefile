.PHONY: all build

build:
	@go build -o bin/calc ./cmd/calc/main.go

run: build
	@./bin/calc
