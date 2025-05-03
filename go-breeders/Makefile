BINARY_NAME=go-breeders-fe

build:
	go build -o ./dist/${BINARY_NAME} ./cmd/web

run:
	go run ./cmd/web

run-cache:
	go run ./cmd/web -cache=true

.PHONY: run run-cache build
