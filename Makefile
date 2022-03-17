NAME := venstar-manage
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	golint ./...
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

static: vet
	staticcheck ./...
.PHONY: static

check: vet lint static
.PHONY: check

build:
	go build -o build/$(NAME)
.PHONY: build

clean:
	rm -f build/*
.PHONY: clean

tidy:
	go mod tidy
.PHONY: tidy
