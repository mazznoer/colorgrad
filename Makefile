SHELL := /bin/bash

.PHONY: all check test

all: check test

check:
	go build && go vet && gofmt -s -l .

test:
	go test -v -coverprofile coverage.out && go tool cover -html coverage.out -o coverage.html

bench:
	go test -bench .
