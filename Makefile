all: build-lint

build:
	go build ./...

build-lint: build lint

test: build
	go test ./... --timeout 30m --count 1 -failfast

install:
	go install ./...

lint:
	golangci-lint run

.PHONY: all build test lint
