all: build-lint

build:
	go build ./...

build-lint: build lint

test: build
	go test ./... --timeout 30m --count 1 -failfast

test-short: build
	go test ./... --short --count 1 -failfast

bench: build
	go test -bench . ./...

install:
	go install ./...

lint:
	golangci-lint run

.PHONY: all build test test-short lint
